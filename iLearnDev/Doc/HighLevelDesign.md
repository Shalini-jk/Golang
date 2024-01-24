# High Level design for the iLearn platform.

There are two component, the server and the client.

The server component is called MyGuru and the client component is called iLearn.

iLearn is expected to manage the whole learning journey for the student. 

Learning is not same as teaching, learning happens by the enquiry and by doing the activities. 

Key challenges in learning any skill are:
- By default, learning journey is not bounded. A student does not know that it does not know, and hence learning journey becomes an undefined journey. An undefined journey never completes.
- The progress on learning is not measurable, and it makes it difficult to sustain. The learning journey pauses, gets disrupted more often.
- As the goal of most of the learning journeys is skill development, skill development happens only by performing activities.  
- It is important to get the personalized feedback on all the activities performed.

# Brand name
iLearn, it will be a page under the www.ambersoft.in 

# iLearn solution
iLearn plans to address all these problems, and develop an AI powered learning platform.

## High Level Features to solve learning problem
1. Bounded content organized as topics.
2. For each topic, a set of kbits to explain the topic. 
3. A set of problems to make sure that the topic is learned. 
4. The nature of activities needs to be start with very simple activities and then gradually activities should become more complex. 
5. AI assisted personalized feedback and rating system for each activity.
6. The Bootcamp mode of group learning, to learn by reviewing peer solutions. 

# User stories
1. User Management: 
    - User is added to be a part of a program by admin/mentors on portal. 
    - User authenticates using google and logins on the portal.
    - User can be part of multiple skill development programs.
2. Learning Content:
    - Each skill development journey is divided in multiple milestones
    - Each milestone has a set of topics
    - Each topic has a set of kbits and a set of problems to solve
3. Learning journey:
    - User is added to the skill development program on the portal
    - User creates personal git repo for this program 
    - User checks out the repo and configures iLearn for this program in the repo
    - User gets all the milestones as folders
    - For the first milestone, user gets the list of topics defined at that point again as folders
    - User start learning each topic. User gets the kbits, and problems for the topic.
    - User can get the list of all the milestones and status of all topics in a milestone, both on the CLI as well as on the portal
    - User solves a problem and submits the solution (file or directory). 
        - iLearn uses AI tools to get a structured feedback and saves the feedback in a file.
        - Both the solutions as well as AI feedback is stored on the server
        - AI generated Rating for each problem is saved separately to track the learning progress.
    - User marks learning completed for a specific topic. 
    - Offline PR view process does not seem to work so far, so peer review will be done at the team level as part of the bootcamp team activities. 
    - User can choose to revise a specific topic, if content for the topic has changed, then user will receive the new content. 
    - User wants to prepare for an exam, or want to raise the bar, then based on the learning history, system can provide specific suggestion about specific topic to revise and solve problems.
    - User's rating at each problem level is a very valuable information to prepare the student for any competitive learning goal, be preparing for an interview, preparing for an exam or competition. This aspect is just a hypothesis at this point, needs to be further validated.
      
# Use of AI
- AI is used by the mentor to create content
- AI is used by the student to get a feedback on the solution 

# operations/APIs to be implemented

## Role management
System has fixed set of following roles:
- Admin: manages programs, students, mentors
- Mentor: manages content and boot camp execution
- Student: manages learning journey

## Learning related entities
- Skill: 
- Batch: With boot camp, we should move to a model of batch. 

## Admin operations
- Admin adds user to the portal 
- Admin assigns a role to the user
- Admin add user to a batch

## Mentor content management operations
- Mentor adding/updating/disabling a skill. Any skill with active students can never get disabled. 
- Mentor adding/updating/deleting milestones for a skill. A milestone with associated content can't be deleted.
- Mentor adding/updating/deleting topics for a milestone. Update also include changing the order for topics. 
- Mentor adding/updating/deleting kbits for a topic.
- Mentor adding/updating/deleting problems for a topic.

## User operations
- sign in / forgot password / authentication

## Student operations
- iLearn setup for a skill
- Get the milestones/topics
- Start learning a topic
- Submit solutions to one or more problems, it result in getting feedback on that topic from chatpGPT and recording feedback in the backend of the student.
- Complete a topic, it requires submitting a summary file, and getting feedback on that summary file.
- Revise a topic

# System Design
For programming skill, expecting following tools to be used:
- Git to keep track of all the content and activities, as well as get peer feedback.
- VScode as the code editor. But no integration needed with vscode at this point.
- AI tools like ChatGPT. Let's start with ChatGPT first and then expand to other AI tools in future. 
- iLearn tool should be main interface for the learning.

Does it make sense to keep the whole content completely git based, does it add any value in adding the data or metadata in the tool.
- For now taking a position to keep all the content in the git and use pathname in the metadata
- For each skill, there is a separate dir for each skill in the iLearn dir.
- For each skill, folders for each milestone, topics and problems
- This seems to be working fine, no need to change that for now.
- still implement the iLearn commands to generate feedback with right prompt engineering
- Also provide the summary of the progress to the student
- Tracking starting and completion of each topic
- 

## CLI / Operation spec
iLearn as client program. 

//List the support skills. ilearn will have lower case l only. No upper case in the command name.
ilearn skill list

//setup a specific skill for the user. Take the user git repo as an input. This operation will figure out the iLearn mentor git repo, checks it out to a well defined location under iLearn directory under a .mentorRepo. And also setup the user git Repo under iLearn/Skill directory. 
//for this user directory, setup the environment ILEARNSKILL to the selected skill.
ilearn skill setup -s skillName -u userEmail -r userGitRepoForSkill


//list all the milestone for the currently selected skill, if skill is not set for the current directory, then ask user to run ilearn skill setup first
ilearn milestone list

//configure topics for a milestone. It will copy all the topics for the milestone from mentor repo to the user repo
ilearn milestone setup -m mileStone

//user can update the list of topics for a milestone. It update kbits for all the topics for that milestone too
ilearn milestone update -m mileStone

//there is no separate topic list operation. User can find list of topics from directories under milestones
//topic setup will get all the kbits for the topic, let's not get all the problems at this point
ilearn topic setup -t <topic_name> -m <milestone_name>

# class management related operations

#mentor starts a class, this operation is available only to mentor, this operation returns a class ID, input for this operation is batchid
myguru class start -b <batchid>

#need add, update, disable operations also for mcq questions
#keeping priority low for mcq 
#generate list of mcq questions for specific topic
myguru mcq list -s <skill_name> -t <topic_name> 

#post operation gives the question to the class, next ilearn mcq <qud> would have shown this question
#there can be only one pending mcq
myguru mcq post -c <class_id> -q <question_id>

#Get the number of response for the mcq 
mygyru mcq responses -c <class_id> -q <question_id>

#close the question, no new submission will be accepted, leaderboard will be updated.
myguru mcq close -c <class_id> -q <question_id>

#show the leaderboard for the class
myguru leaderboard -c <class_id>

#show the summary leaderboard for the class with top 5 entries
myguru leaderboard -c <class_id> -s

#join the class, mentor should have started the class and that generates the class code 
ilearn class join <classcode> -d <daily_update>

#give mcq question, gets the question and choices and gets a prompts to give ht
ilearn mcq <q#>

#get the leaderboard
ilearn leaderboard -c <class_id>

#get summary leaderboard, top 5 and me
ilearn leaderboard -c <class_id> -s

#overall leaderboard across all topics for the batch
ilearn leaderboard -b <batch_id>

#Get summary leaderboard across all topics for the batch and me
ilearn leaderboard -b <batch_id> -s

# revision activity

#start a revision activity, this will mark the time stamp of starting of such activity
#only one topic revision can be in start state for that student
ilearn revision start -t <topic_name> 

#complete the revision, provide the notes to summaries the understanding, tool will send these notes to AI and get feedback on these notes. 
#feedback is saved in notes_fb.md file under topic directory
ilearn revision complete -t <topic_name> -n <notes.md>

# Asking question to my Guru

#Ask doubt clear question to my Guru
#Answer is saved in question_ans.md file in the same directory
ilearn ask -t <topic_name> -q <question.md>

# Problem solving

#Get next problem to solve, it might be a subjective or programming problem. 
#myGuru will look into your past history for the topic and give appropriate problem.
#problem will be saved in your directory in a .md file and name will be displayed
#time to solve problem starts when you receive the problem
ilearn problem get -t <topic_name> 
#submit the solution to the problem, myGuru will generate feedback using AI and save the feedback in problem_fb.md file
ilearn problem solve -t <topic_name> -p <problem_file_name> -s <solution_file_or_dirname>

# File as main interface
Files are the main interface to share content, problems, feedback. 

# directory format for setting up the git repo
# Here is the suggested directory structure, this is created at the time of skill setup
- iLearn
    - Golang (name of skill)
        - .MyGuru (directory containing the main content provided by MyGuru)
             - BasicTopics
                1_intro
        - MyRepo User git repo
            - BasicTopics
                1_intro

# Server side design
- Let's start with the SQLite3 and GORM
- User management
    - user authentication: can we use the google auth, I don't want to manage password
    - user to skill mappings, single user can setup multiple skills
        - myGuru will add users to the system and also add the skill for which they have registered, skill setup will validate both useremail and skill
        - each user will have multiple roles, there are two roles in the system, mentor and student
    - there will be two client programs
        - iLearn: for the student
        - iMentor: for the mentors
    - there will be a server program called myGuru
- There are two framework to support REST APIs, these are Gorilla and Gin. Gorilla might be simple to use, for now I am starting with Gorilla. But if needed, in future can switch to Gin too.

# Google Auth setup
Created following project:
https://console.cloud.google.com/home/dashboard?project=ilearn-411009
Under credentials, you can see iLearnClient for the OAuth 2.0 client ID.

# REST API design

Resources to manage:
1. Users: Only mentor can add, delete users. We should allow user to get and update the user details. 
2. Skills: Only mentor can add, update and delete skills. User can get the list of skills.
3. Users/{userID}/skills/: Get all the skills for the user, only mentor should map a specific skill to a user
4. Users/{userID}/skills/{skillID}/milestones: Manage milestones 
5. Users/{userID}/skills/{skillID}/milestones/{milestoneID}/topics: manage topics

# Thoughts
structure of milestones and topics is nice, but it seems that learner are struggling directly jumping to the problems.
Most of the problems seem to be too difficult for them. 
Expectation of the students is a gradual ramp up of activities.

There are following learning related activities 
- For each topic, can we really have a bounded list of concepts. As we engage in the learning, these concepts might evolve.
- It will be difficult to track each concept separately.
- User should be allowed to suggest new kbits for topic or update to each topic. For now, there is no interface for this, let
s mentor discover this and keep the kbits updated.
- Rating should be tracked for each topic
- There are only following types of activities:
    - Providing daily update, it should also be structured
    - Attending class session, always a broadcast, can have notes/video recording, with info about discussion on specific topics. So class should always have a reference to a set of topics. Should have a notion of attendees, and an AI generated summary should be provided for each meeting.
        - class should have MCQ and leader board score to make sure that students are engaging as well as they are able to understand
        - Guru takes the feedback from MCQ to adjust the delivery
    - Team sync, team of 4-6 student can have sync, and this should also have reference to topic getting discussed. AI generated summary should be provided for this too.
    - Individual learning activity
        - User records what kind of activity 
        - Reading: student is revising the content provided, writing some sample code etc. Notes should get created for reading/revising activity too.
        - Asking: Student need clarification from My Guru, asks questions. These should be recorded. This can be an independent activity, or it can be done in parallel with reading activity.
        - Problem Solving: Subjective as well as programming. These results in student getting grades. And complexity of problem is gradually increased, all the problems are not given at the same time.

- Here is iLearn scheme of things
  - There are a set of students in a class for a skill
  - There is a mentor for the class
  - Each student configures the skill with ilearn
  - The mentor starts the class and give the class code to the students
  - Each student joins the class with the class code
  - Class content is provide in term of .md file under the appropriate topic
  - Mentor releases specific MCQ question, waits for all the responses
  - Mentor can show / release the leader board for the class
  - Student can start a reading session or asking session. Each session start and end times are recorded.
  - Student can start a problem solving session for a topic. Only one problem is shown to the user at a time. Once problem is submitted, detailed AI generated feedback as well as correct solution is shown. User rating for the each topic is determined by the complexity of the problem, time taken. It is alway generated by AI.
  - For each team meeting, team lead for the day calls for the meeting, shares the meeting code with the students, students joins the meetings, team lead posts the meeting notes.




2. We should allow the user to suggest new kbit for the topic. 

