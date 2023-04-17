<!--
-
-
	type ``` ctr+shift+v ``` to open markdown preview in VSC (to see the pretty version)
-
-
!-->

# Github Common Commands
may or may not work. im not an expert  
All git commands done through git bash (or any terminal prob) in a folder containing all docs: \<New Folder>

## Contents
- [Angular Setup](#angular-setup)
- [Starting up Angular in VSC](#starting-angular)
- [running GO in VSC](#running-in-vsc)
- [Pull from repo](#pull-from-repo)
- [Push to Repo (from a pull)](#push-from-a-pull)



# Angular Setup
assuming you have not pulled any go files from the repo -- just starting fresh.  
after creating an angular project using 
```
npm new <newAng> (name it frontend or templates or smth) 
```

in your file explorer
- create a new folder
- put \<newAng> inside that new folder.  
- so now you are in new folder. and \<newAng> folder is the only folder visible
(you may delete the .git folder inside \<newAng>. idk if its necessary but i did)  
  
git bash from the new folder
``` 
git init
git remote add origin <repo url>
git pull origin main 
git add * 
```
there is already a .gitignore file made automatically by angular. but make sure nothing from node_modules gets added. (you will notice because git add will take a veery long time. just press ctrl + c to stop .. ask if it happens but it shouldnt)
```
git commit -m "commit msg"
git push --set-upstream origin main
git push 
```

# Starting Angular
Create a new folder (prob the easiest way to make sure everything uploads clean)
git bash from the new folder
``` 
git init
git remote add origin <repo url>
git pull origin main 
```
Use command prompt (vscode terminal didn't work for me)
``` 
cd <the new folder>
cd frontend
npm install 
ng serve

you can open your browser to localhost:9000. 
localhost:4200 is just angular. 9000 is go and angular.


# Running in VSC
after pulling go files from repo:  

make sure go.mod is in the same dir as main.go
## to run the go program from VSC:
```
open <New Folder>  
```
type in vsc terminal:
```
go run main.go
```

## if go.mod does not exist: 
type in vsc terminal:
```
go mod init go
```
make sure to type this in the same dir as main.go  
go.mod should appear in the same folder as main.go

## if go.mod is in the wrong dir:
	delete go.mod
	create go.mod using same procedure from above



# Pull from Repo 
Pulling changes made to the repo and applying them to local files.  
Pull before making changes and pushing.
## For the First Time
type in git bash from \<New Folder>:
```
git init
git pull <repo url>
```

## Normally
type in git bash from \<New Folder>:
```
git pull
```


# Push from a Pull
Pushing local changes to the repo.  
(make sure your local files are up-to-date before making changes and attempting to push) 
## For the First Time (after pulling for the first time)
type in git bash from \<New Folder>:
```
git remote add origin <url>
git branch -M main
git add *
git commit -m "commit msg"
git push --set-upstream origin main
git push
```
## Normally
type in git bash from \<New Folder>:
```
git add *
git commit -m "commit msg"
git push
```
