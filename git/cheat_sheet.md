git commit --amend --reset-author             #change code author



git reflog --date=local | grep <branchname>   #query source branch name



git remote add XXX http://A.git               #add other's code rep



git fetch --all                               #get new commit



git reset --hard origin/xxx                   #pull -f



=====================================================================================================



git pull upstream huawei/master               # pull new commits from old code hub（for fork）



git branch|grep 'r20a'|xargs git branch -D    # delete branchs named r20a



git reset --soft HEAD@{1}                     # rollback git commit --amend



git config --list                             # check current git configuration



git branch -m oldbranch newbranch             # rename your branch



=====================================================================================================



git checkout -b branch1                 # creat a new branch, named branch1



git checkout branch1                    # move on the branch1



git checkout -f                         # if you deleted some file, use it to recover



git checkout filename                   # recover the unsubmited file


git checkout .                          # recover all unsubmited files



git status                              # show current git status



git add filename                        # add a new file to current version



git add .                               # add all new files



git commit -a                           # commit all files



git commit                              # saved changes are called commits



git commit -m "description"             # each commit has an associated commit message. 

                                          ESC+:+wq  # save decsription



git commit -am "new added files"        # -a and -m



git push --set-upstream origin branch1  # push local files to remote branch1

                                          --set-upstream is for 1st time



git merge branch1                       # merging means bringing changes together



=====================================================================================================



git log                                 # show all commits



git log -1                              # show the lastest commit



git log --name-status                   # files modification status



git log --graph --pretty=oneline        # show all commits graphically



git log --oneline -5                    # show 5 commits abbreviated to one line



git commit --amend                      # keep one commit description automatically



git commit -amend -no-edit              # stole one's code ^_^



git rebase -i commitSHA                 # rebase latest 2 commits

                                          commitSHA is not necessary（-pick, -s）



git rebase --continue                   # rebase latest 2 commits automatically



git rebase -i HEAD~n                    # rebase latest n commits



git push origin branch1                 # push local files to remote



git push -f origin branch1              # force push local files to remote



=====================================================================================================



git pull = git fetch + git merge



git fetch origin master:temp            # creat a new branch locally named temp,

                                          download repositories from master into temp


git diff temp                           # compare codes with local codes



git diff branch1 origin/branch1         # compare codes with remote branch



git diff d72341 b234a2                  # compare codes with 2 commits


git merge temp                          # merge temp into local branch



git merge branch1 master                # merge source branch1 into master

git branch -d temp                      # delete the temp branch


git cherry-pick commitSHA               # pick one commit to current branch



git cherry-pick branch1                 # pick the top one commit from branch1



git cherry-pick branch1~n               # pick the top n-th commit from branch1


=====================================================================================================



git remote                                    # list remote hostname



git rm file                                   # remove your file



git mv file folder/                           # move your file to a folder, i.e. renamed the file



git reset --hard d72723                       # back to some commit + checkout



git reflog                                    # git command history log



git rebase --abort                            # abort rebasing



git rebase masterbranchname                   # get the latest commits in master branch



=====================================================================================================



git checkout -b branchname commitSHA          # use some old commit to make a branch for patches



git reset HEAD                                # come back to the time before "add"



git reset HEAD XXX.c                          # pick a c file come back to the time before "add"


=====================================================================================================



git documentation：https://git-scm.com/docs
