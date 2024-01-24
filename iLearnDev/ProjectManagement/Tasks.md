# Use this doc to track all the pending tasks and track the daily progress for each task.

## WIP Tasks
- High Level design: Taken a shot at defining all the operations at a high level
- ChatGPT API learning: Starting. Tracking details doc/ChatGPTAPILearning.md
- 

## Pending Tasks 
- need to handle the case where command/subcommands are not valid, it should print usage message in those cases.
- implement my guru server with DB to start supporting REST APIs between client and server without any authentication
- implement end to end client/server operations
- start adding unit tests
- implement revision operations
- implement ask operations
- implement problem operations
- implement class operations
- implement mcq operation
- implement leader board
- implement skill server side operation
- User authentication using OAuth 2.0
- Set up the MyGuru repo as read only repo for the users so that they can't checkin there
- config file to manage various parameter, for now ILEARNSKILL is manually set

# Tasks for team
1. Unit tests: skill Shalini
2. Unit tests: milestone Raj
3. Unit tests: topic Chitra
4. Unit tests for fileops: Amit

## Completed Tasks
- Need to create initial dir structure and checkout the git repo
- implement skill client side operations
- implement logging system, don't use just the normal Println
- implement milestone client side operations
- fix .gitignore, currently it is not tracking newly added files. Goal was not to include any binary files.
- implement topics client side operations

# test cases
iLearn skill list
iLearn skill setup -s Golang -u anurag@ambersoft.in -l https://github.com/anurag-kqi/LearningGolang.git
iLearn topic setup -m BasicTopics -t 3_basic_data_types

setup tests
1. iLearn skill setup -s Golang -u anurag@ambersoft.in -l https://github.com/anurag-kqi/LearningGolang.git
   worked, created ~/iLearn/Golang/myRepo dir and cloned the repo there
2. running setup again fails as target directory is there
3. running setup with invalid gitRepo URL fails as expected 


## Daily Notes
16th Jan, 24:
- start with the myGuru, a DB application
- It is first time I am dealing with Gorm and DB, able to add Gitrepo in the join table, in the sample code. Let me cleanup the code now. 
  Name of the tables are derived from the name of structure, I was expecting that there will be separate join table and I need to create a new table for adding the GitRepo field, but that is not true. I am able to add the field in the join table by using appropriate names.
- names of fields needs to be chosen carefully.
- Added DB support to add skill, user and update user,skill mapping.
- Need to create REST APIs for this now


12 Jan, 24:
- Cleaned up iLearn repo, now let's put the content in .MyGuru folder where iLearn repo has been checked out.
- fix .gitignore, currently it is not tracking newly added files. Goal was not to include any binary files. Done
- implement topics client side operations

11th Jan, 24:
- milestone functionality: Done 

10th Jan, 24:
- let's implement skill functionality fully
  - let's move the skill related code to skill.go: Done
  - let's work on the skill setup operation: Done
    - create the directory and clone the user repo
  - Added .gitignore to ignore adding of binary files
  - Set up logging support: Done
    - let's use the logrus so that we can have different log level.



9th Jan, 24:
- Let's iterate on the cli interface once more to make sure that cli is well defined
- Expanded the scope of the problem reasonably well, define the operations
- let's start implementing the first set of operations around setting up the course and milestone, topics etc.
- let start working on the milestone commands
2nd Jan, 24:
