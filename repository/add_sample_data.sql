insert into users(email, name, role, department, bio, photo_url, contact, is_deleted, onboarded, github_url)
VALUES ('tushar@email.com', 'Tushar Paliwal', 'Intern', 'Risk', 'Tech Enthusiast',
        'https://avatars3.githubusercontent.com/u/55799457?s=460&v=4', '@tushar', false, true, 'https://github.com/tushar'),
       ('arjun@email.com', 'Arjun Ramachandran', 'Intern', 'Risk', 'Tech Enthusiast',
        'https://avatars3.githubusercontent.com/u/55799457?s=460&v=4', '@arjun', false, true, 'https://github.com/tushar'),
       ('prashant@email.com', 'Prashant Agarwal', 'Intern', 'Risk', 'Data science Enthusiast',
        'https://avatars3.githubusercontent.com/u/55799457?s=460&v=4', '@prashant', false, true, 'https://github.com/tushar');

insert into globalskills(created_by, value, time_created)
VALUES (1, 'nodejs', '2019-06-22 19:10:25-07'),
       (1, 'spring', '2019-06-22 19:10:25-07'),
       (2, 'react', '2019-07-22 19:10:25-07'),
       (2, 'golang', '2019-07-22 19:10:25-07'),
       (3, 'tableau', '2019-07-23 19:10:25-07'),
       (3, 'powerbi', '2019-07-23 19:10:25-07'),
       (3, 'spark', '2019-07-23 19:10:25-07');

insert into userskills(user_id, skill_id, time_created, time_updated, is_deleted)
VALUES (1, 1, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (1, 2, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (2, 1, '2019-07-22 19:10:25-07', '2019-07-22 19:10:25-07', false),
       (2, 3, '2019-07-22 19:10:25-07', '2019-07-22 19:10:25-07', false),
       (2, 4, '2019-07-22 19:10:25-07', '2019-07-22 19:10:25-07', false),
       (3, 5, '2019-07-23 19:10:25-07', '2019-07-23 19:10:25-07', false),
       (3, 6, '2019-07-23 19:10:25-07', '2019-07-23 19:10:25-07', false),
       (3, 7, '2019-07-23 19:10:25-07', '2019-07-23 19:10:25-07', false);

INSERT INTO JOBS (created_by, title, description, difficulty, status, time_created, time_updated, is_deleted)
VALUES (1,
        'This is a sample job the needs help with react, spring and nodejs',
        'In the limitation flow when selecting the "explain activity" lifting requirement the calendar does not allow selection of the current date. This seems to be a new feature that was implemented in a recent push and it is important for the fraud dept to be able to select the current date since they review accounts real time.',
        'Intermediate',
        'open',
        '2019-06-22 19:10:25-07',
        '2019-06-22 19:10:25-07',
        'FALSE'),
       (1,
        'Create authentication flow for Innersource project',
        'Innersource is a project to let people find cool projects to work on within their organisation. However, there''s no authentication for now and we require that to be done before launching',
        'Intermediate',
        'ongoing',
        '2019-07-22 19:10:25-07',
        '2019-07-22 19:10:25-07',
        'FALSE'),
       (2, 'This is another job that needs help with powerbi, tableau and spark',
        'In the limitation flow when selecting the "explain activity" lifting requirement the calendar does not allow selection of the current date. This seems to be a new feature that was implemented in a recent push and it is important for the fraud dept to be able to select the current date since they review accounts real time.',
        'Intermediate',
        'open',
        '2019-08-22 19:10:25-07',
        '2019-08-22 19:10:25-07',
        'FALSE');

INSERT INTO milestones (job_id, title, description, duration, resolution, status, time_created, time_updated,
                        is_deleted)
VALUES (1,
        'Sample milestone 1 for job 1',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. ',
        '1 week',
        'github pull request,Pass unit tests',
        'open',
        '2019-06-22 19:10:25-07',
        '2019-06-22 19:10:25-07',
        'FALSE'),
       (1,
        'Sample milestone 2 for job 1',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',
        '1 week',
        'github pull request,Pass unit tests',
        'open',
        '2019-06-22 19:10:25-07',
        '2019-06-22 19:10:25-07',
        'FALSE'),
       (2,
        'Setup CHI router to handle middleware',
        'CHI router is the go-to router for production softwares. Setup chi router so that it handles existing routes instead of the basic serv-mux.',
        '1 week',
        'github pull request,Pass unit tests',
        'ongoing',
        '2019-07-22 19:10:25-07',
        '2019-07-22 19:10:25-07',
        'FALSE'),
       (2, 'Create auth middleware to handle JWT authentication',
        'Setup JWT authentication based on one of the popular jwt libs (your choice) but do give a reason for your choice. Ensure short expiry time',
        '1 week',
        'github pull request,Pass unit tests',
        'ongoing',
        '2019-07-22 19:10:25-07',
        '2019-07-22 19:10:25-07',
        'FALSE'),
       (2, 'Setup a way to revoke and blacklist tokens',
        'To increase the security of the application, setup token revocation of blacklisting in the auth middleware',
        '1 week',
        'github pull request,Pass unit tests',
        'ongoing',
        '2019-07-22 19:10:25-07',
        '2019-07-22 19:10:25-07',
        'FALSE'),
       (3,
        'This a sample milestone 1 of job 3',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',
        '1 week',
        'github pull request,Pass unit tests',
        'open',
        '2019-08-22 19:10:25-07',
        '2019-08-22 19:10:25-07',
        'FALSE'),
       (3,
        'This a sample milestone 2 of job 3',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',
        '1 week',
        'github pull request,Pass unit tests',
        'open',
        '2019-08-22 19:10:25-07',
        '2019-08-22 19:10:25-07',
        'FALSE'),
       (3,
        'This a sample milestone 3 of job 3',
        'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',
        '1 week',
        'github pull request,Pass unit tests',
        'open',
        '2019-08-22 19:10:25-07',
        '2019-08-22 19:10:25-07',
        'FALSE');

insert into milestoneskills(milestone_id, skill_id, time_created, time_updated, is_deleted)
values (1, 2, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (1, 1, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (2, 1, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (2, 3, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (3, 3, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (4, 3, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (5, 2, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (3, 4, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (4, 4, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (5, 4, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (5, 3, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (6, 5, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (6, 6, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (7, 7, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false),
       (8, 7, '2019-06-22 19:10:25-07', '2019-06-22 19:10:25-07', false)
;

insert into discussions(job_id, created_by, content, time_created, time_updated, is_deleted)
values (1, 1, 'such job, many discussions', '2020-07-22 19:10:25-07', '2020-07-22 19:10:25-07', false),
       (1, 2, 'Doesn''t look like anyone has picked this up yet. I will do it', '2020-07-22 19:10:25-07',
        '2020-07-22 19:10:25-07', false),
       (2, 2, 'Any specific reason for chi router rather than gorilla mux?', '2020-07-22 19:10:25-07',
        '2020-07-22 19:10:25-07', false),
       (2, 1, 'Good question, we should compare the two', '2020-07-22 19:10:25-07', '2020-07-23 19:10:25-07', false),
       (2, 1, 'Should go with gorrilla because it has more github stars', '2020-07-22 19:10:25-07',
        '2020-07-23 19:12:25-07', true),
       (2, 1, 'We should go with chi because they already have a decent jwt auth middleware', '2020-07-22 19:13:25-07',
        '2020-07-23 19:13:25-07', false);

insert into applications(milestone_id, applicant_id, status, note, time_created, time_updated)
values (1, 1, 'rejected', 'Sorry, you don''t seem to have sufficient experience', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (2, 1, 'rejected', 'Sorry, you don''t seem to have sufficient experience', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (3, 2, 'accepted', 'Contact me @ lmao@email.com', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (4, 2, 'accepted', 'Contact me @ lmao@email.com', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (5, 2, 'accepted', 'Contact me @ lmao@email.com', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (6, 3, 'pending', '', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (7, 3, 'pending', '', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (8, 3, 'pending', '', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (6, 1, 'withdrawn', '', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (7, 1, 'withdrawn', '', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07'),
       (8, 1, 'withdrawn', '', '2020-06-22 19:10:25-07',
        '2020-06-23 19:10:25-07');

