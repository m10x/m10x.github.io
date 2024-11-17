package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {

	strThemeURL := "https://github.com/rhazdon/hugo-theme-hello-friend-ng.git"
	strTheme := "hello-friend-ng"
	strBlogPath := "m10x-blog"
	strThemeFolderPath := strBlogPath + "/themes/"
	strThemePath := strThemeFolderPath + strTheme

	fmt.Println("----- 0) Updating theme")

	os.RemoveAll(strThemeFolderPath)

	out, err := exec.Command("git", "clone", strThemeURL, strThemePath).Output()
	if err != nil {
		log.Fatal("git clone: ", err)
	}
	fmt.Println("Git Clone:", string(out))

	fmt.Println("----- 0) Finished")

	fmt.Println("----- 1) Removing everything in public but dont remove CNAME")

	dir, err := ioutil.ReadDir(strBlogPath + "/public")
	for _, d := range dir {
		if d.Name() != "CNAME" {
			os.RemoveAll(path.Join([]string{strBlogPath + "/public", d.Name()}...))
		}
	}

	fmt.Println("----- 1) Finished")

	fmt.Println("----- 2) Running hugo")

	out, err = exec.Command(strBlogPath+"/hugo", "-t", strTheme).Output()
	if err != nil {
		log.Fatal("hugo: ", err)
	}
	fmt.Println("Hugo:", string(out))

	fmt.Println("----- 2) Finished")

	fmt.Println("----- 3) Pushing changes")

	/* max@max-pc MINGW64 ~/Documents/git/blog/m10x-blog/public (master)
	$ git add *
	fatal: in unpopulated submodule 'm10x-blog/public'


	git submodule add git@github.com:m10x/m10x.github.io.git public
	git submodule add git@github.com:rhazdon/hugo-theme-hello-friend-ng.git themes/hello-friend-ng
	*/

	out, err = exec.Command("git", "-C", strBlogPath+"public", "add", "-A").Output()
	if err != nil {
		log.Fatal("git add: ", err)
	}
	fmt.Println("Git Add:", string(out))

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a commit message: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("enter message: ", err)
	}

	out, err = exec.Command("git", "-C", strBlogPath+"public", "commit", "-m", text).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Git Commit:", string(out))

	out, err = exec.Command("git", "-C", strBlogPath+"public", "push").Output()
	if err != nil {
		log.Fatal("git push: ", (err))
	}
	fmt.Println("Git Push:", string(out))

	fmt.Println("----- 3) Finished")
}
