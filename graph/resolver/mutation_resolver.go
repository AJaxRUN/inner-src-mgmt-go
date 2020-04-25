package resolver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	gqlmodel "github.com/cassini-Inner/inner-src-mgmt-go/graph/model"
	"github.com/cassini-Inner/inner-src-mgmt-go/middleware"
	dbmodel "github.com/cassini-Inner/inner-src-mgmt-go/postgres/model"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"strings"
	"time"
)

var (
	ErrUserNotAuthenticated           = errors.New("unauthorized request")
	ErrUserNotOwner                   = errors.New("current user is not owner of this entity, and hence cannot modify it")
	ErrNoEntityMatchingId             = errors.New("no entity found that matches given id")
	ErrOwnerApplyToOwnJob             = errors.New("owner cannot apply to their job")
	ErrApplicationWithdrawnOrRejected = errors.New("owner cannot modify applications with withdrawn status")
	ErrInvalidNewApplicationState     = errors.New("owner cannot move application status to withdrawn or pending")
	ErrJobAlreadyCompleted            = errors.New("job is already completed")
)

func (r *mutationResolver) MarkMilestoneCompleted(ctx context.Context, milestoneID string) (*gqlmodel.Milestone, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MarkJobCompleted(ctx context.Context, jobID string) (*gqlmodel.Job, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProfile(ctx context.Context, updatedUserDetails *gqlmodel.UpdateUserInput) (*gqlmodel.User, error) {
	currentRequestUser, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUserNotAuthenticated
	}

	updatedUser, err := r.UsersRepo.UpdateUser(currentRequestUser, updatedUserDetails)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *mutationResolver) CreateJob(ctx context.Context, job *gqlmodel.CreateJobInput) (*gqlmodel.Job, error) {
	if len(job.Desc) < 5 {
		return nil, errors.New("description not long enough")
	}
	if len(job.Title) < 5 {
		return nil, errors.New("title not long enough")
	}
	if len(job.Difficulty) == 5 {
		return nil, errors.New("diff not long enough")
	}
	if len(job.Milestones) == 0 {
		return nil, errors.New("just must have at least one milestone")
	}
	for _, milestone := range job.Milestones {
		if len(milestone.Skills) == 0 {
			return nil, errors.New("milestone must have at least one skill")
		}
	}
	user, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	newJob, err := r.JobsRepo.CreateJob(ctx, job, user)
	if err != nil {
		return nil, err
	}
	var gqlJob gqlmodel.Job
	gqlJob.MapDbToGql(*newJob)
	return &gqlJob, nil
}

func (r *mutationResolver) UpdateJob(ctx context.Context, job *gqlmodel.UpdateJobInput) (*gqlmodel.Job, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteJob(ctx context.Context, jobID string) (*gqlmodel.Job, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddCommentToJob(ctx context.Context, comment string, jobID string) (*gqlmodel.Comment, error) {
	user, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	newComment, err := r.DiscussionsRepo.CreateComment(jobID, comment, user.Id)
	if err != nil {
		return nil, err
	}

	var gqlComment gqlmodel.Comment
	gqlComment.MapDbToGql(*newComment)
	return &gqlComment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, comment string) (*gqlmodel.Comment, error) {
	user, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUserNotAuthenticated
	}

	existingDiscussion, err := r.DiscussionsRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoEntityMatchingId
		}
		return nil, err
	}
	if existingDiscussion.CreatedBy != user.Id {
		return nil, ErrUserNotOwner
	}

	updatedDiscussion, err := r.DiscussionsRepo.UpdateComment(id, comment)
	if err != nil {
		return nil, err
	}
	var gqlUpdatedDiscussion gqlmodel.Comment
	gqlUpdatedDiscussion.MapDbToGql(*updatedDiscussion)
	return &gqlUpdatedDiscussion, nil
}

func (r *mutationResolver) DeleteCommment(ctx context.Context, id string) (*gqlmodel.Comment, error) {
	if id == "" {
		return nil, errors.New("invalid comment id")
	}

	user, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUserNotAuthenticated
	}

	existingDiscussion, err := r.DiscussionsRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoEntityMatchingId
		}
		return nil, err
	}
	if existingDiscussion.CreatedBy != user.Id {
		return nil, ErrUserNotOwner
	}
	err = r.DiscussionsRepo.DeleteComment(id)
	if err != nil {
		return nil, err
	}
	existingDiscussion.IsDeleted = true
	var gqlComment gqlmodel.Comment
	gqlComment.MapDbToGql(*existingDiscussion)
	return &gqlComment, nil
}

func (r *mutationResolver) CreateJobApplication(ctx context.Context, jobID string) ([]*gqlmodel.Application, error) {
	var result []*gqlmodel.Application

	user, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUserNotAuthenticated
	}

	job, err := r.JobsRepo.GetById(jobID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoEntityMatchingId
		}
		return nil, err
	}

	if job.Status == "completed" {
		return nil, ErrJobAlreadyCompleted
	}

	if job.CreatedBy == user.Id {
		return nil, ErrOwnerApplyToOwnJob
	}

	milestones, err := r.MilestonesRepo.GetByJobId(jobID)
	if err != nil {
		return nil, err
	}
	applications, err := r.ApplicationsRepo.CreateApplication(milestones, user.Id, ctx)
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		var gqlApplication gqlmodel.Application
		gqlApplication.MapDbToGql(application)
		result = append(result, &gqlApplication)
	}

	return result, nil
}

func (r *mutationResolver) DeleteJobApplication(ctx context.Context, jobID string) ([]*gqlmodel.Application, error) {
	user, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUserNotAuthenticated
	}

	jobMilestones, err := r.MilestonesRepo.GetByJobId(jobID)
	if err != nil {
		return nil, err
	}

	applications, err := r.ApplicationsRepo.SetApplicationStatusForUserAndJob(user.Id, jobID, jobMilestones, dbmodel.ApplicationStatusWithdrawn, nil)
	if err != nil {
		return nil, err
	}

	var result []*gqlmodel.Application

	for _, application := range applications {
		var temp gqlmodel.Application
		temp.MapDbToGql(application)
		result = append(result, &temp)
	}
	return result, nil
}

func (r *mutationResolver) UpdateJobApplication(ctx context.Context, applicantID string, jobID string, status *gqlmodel.ApplicationStatus, note *string) (result []*gqlmodel.Application, err error) {
	// since this end point can only be user by job owner,
	// they can only modify job status from pending to accepted or pending
	currentStatus, err := r.ApplicationsRepo.GetApplicationStatusForUserAndJob(applicantID, jobID)
	if err != nil {
		return nil, err
	}

	// check if the currently authenticate user is the owner of the job
	currentUser, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, err
	}
	currentJob, err := r.JobsRepo.GetById(jobID)
	if err != nil {
		return nil, err
	}
	if currentJob.CreatedBy != currentUser.Id {
		return nil, ErrUserNotOwner
	}

	// owner cannot modify the status of application what was withdrawn by applicant
	// owner can only move an application from p
	// - pending->accepted, pending->rejected, accepted->rejected
	if currentStatus == "withdrawn" {
		return nil, ErrApplicationWithdrawnOrRejected
	}
	// owner cannot move the application from pending or withdrawn state to any new state
	if status.String() == "PENDING" || status.String() == "WITHDRAWN" {
		return nil, ErrInvalidNewApplicationState
	}

	milestones, err := r.MilestonesRepo.GetByJobId(jobID)
	if err != nil {
		return nil, err
	}

	updateResult, err := r.ApplicationsRepo.SetApplicationStatusForUserAndJob(applicantID, jobID, milestones, strings.ToLower(status.String()), note)

	if err != nil {
		return nil, err
	}

	for _, application := range updateResult {
		var gqlApplication gqlmodel.Application
		gqlApplication.MapDbToGql(application)
		result = append(result, &gqlApplication)
	}
	return result, nil

}

func (r *mutationResolver) Authenticate(ctx context.Context, githubCode string) (*gqlmodel.UserAuthenticationPayload, error) {
	// authenticate the user with github and store them in db
	user, err := r.UsersRepo.AuthenticateAndGetUser(githubCode)
	if err != nil {
		return nil, err
	}

	// map db user to graphql model
	var resultUser gqlmodel.User
	resultUser.MapDbToGql(*user)
	//generate a token for the user and return
	authToken, err := resultUser.GenerateAccessToken()

	if err != nil {
		log.Println(err)
		return nil, errors.New("something went wrong")
	}
	refreshToken, err := resultUser.GenerateAccessToken()

	if err != nil {
		log.Println(err)
		return nil, errors.New("something went wrong")
	}
	resultPayload := &gqlmodel.UserAuthenticationPayload{
		Profile:      &resultUser,
		Token:        *authToken,
		RefreshToken: *refreshToken,
	}
	return resultPayload, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (*gqlmodel.UserAuthenticationPayload, error) {
	// get the claims for the user
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Printf("error while refreshing refreshToken %v", refreshToken)
		log.Println(err)
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid refreshToken signature")
		}
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("refreshToken is not valid")
	}
	// only refresh the refreshToken if it's expiring in 2 minutes
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > (time.Minute * 2) {
		return nil, errors.New("refreshToken can only be refreshed 2 minutes from expiry time")
	}
	//generate a new refreshToken for the user
	user, err := r.UsersRepo.GetById(claims.Id)
	if err != nil {
		log.Printf("error getting user from claims for user id %v", claims.Id)
		return nil, err
	}
	var gqlUser gqlmodel.User
	gqlUser.MapDbToGql(*user)
	newToken, err := gqlUser.GenerateAccessToken()
	newRefreshToken, err := gqlUser.GenerateRefreshToken()
	if err != nil {
		log.Printf("error generating refreshToken for user %+v", gqlUser)
		return nil, err
	}
	return &gqlmodel.UserAuthenticationPayload{
		Profile:      &gqlUser,
		Token:        *newToken,
		RefreshToken: *newRefreshToken,
	}, nil
}
