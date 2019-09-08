# Portfolio

[![Preview](https://github.com/WesternPine/Portfolio/blob/master/lib/FILES/preview.png?raw=true)](https://westernpine.dev)

^ [Click To Visit Me!](https://westernpine.dev)

## Introduction

My portfolio was my first attempt at a personal/(semi)professional portfolio, or even a website in general. I tried to make this as easy to use and customize as possible, so that others can get a quick head-start on their personal projects. With every website however, comes some customization that needs to be adjusted in the backend of things. All of which will be explained [later](https://github.com/WesternPine/Portfolio#setup).

## Installation (Requires Git):

  - Clone the Portfolio Repository.
  - Look at the beautiful code.
  
## Startup (Requires GoLang, Optional MySQL):

  - Windows:
    - Either run `zStart.bat` or `go run portfolio.go`
  - Linux:
    - Ensure you can run GoLang as a sudo command.
    - Either run `./zStart.sh` or `sudo go run portfolio.go`
  - Other/All OS:
    - run `go run portfolio.go`

## Setup

Setting up the website to work on your server or computer is actually pretty easy, though it does need a little set-up if you want to use all of it's features. To begin, TLS is optional, be sure to read both methods of installation depending on your decision. Additionally, SQL is also optional, though it requires a bit of digging into the code to change some things. Also, instead of disabling MySQL, we will be disabling the entire contact/form functionality of the website to keep things simple. In my instance though, I went with a MySQL Database to store all of the information users submitted to me through my website. Lets start with some general configuration.

### Configuration

If you are not using MySQL, you may skip this part. If you are using MySQL, you might not notice it, but you are actually missing a configuration file. Thus, you will have to create it on your own. Start by creating a new text/json file, and name it: `contactConfig.json` as this is the exact name used in the hosting software written in go. You can change this in the go file. Now, we must populate it. Go ahead and paste the following information in the file you just created.

```
{
  "username": "root",
  "password": "password",
  "ip": "localhost",
  "port": "3306",
  "database": "forms",
  "table": "forms"
}
```

Please make sure the user specified has access to the database via ip added to the network, ip added to user, user added to database.

### MySQL (Setup)

If you decided against MySQL, please skip to the next section. If not, then this is very simple. Get ahold of any MySQL database you're comfortable with. Create a database `forms` (or whatever is set in your `contactConfig.json`).

## Final Setup And Configuration

Now running the site's pages is pretty straight-forward. Go to the `pages.json`file. There you should see some lines of code that specify a url path `/` and a file path `/pages/index.html`. This is where each url path corresponds with the web-page file path. Pretty straight forward. Now dealing with the rest of the website... This is where things can get complicated or easy as pie. Let's start by using the defaults.

### With TLS

Using TLS means you don't have to deal with the website's code. Though, you will need to set-up your own website's certificates. I will not be explaing how to do this here as there are plenty of explinations out there. However, to run TLS on the site, you must have your key named `privkey.pem` and cert named `cert.pem`, and have them both located in the same directory as the `portfolio.go` file. And that's it for TLS! One last change you must make for everything to operate, is by going into the file at `/lib/JS/submit.js` and changing the line url around line 22, to `https://your.domain/formsubmissionhandler` which will save any forms submitted using the server. (Please notice the use of 's' in https. As https is port 443, and http is port 80)

### Without MySQL (With Or Without TLS)

If you decide not to run the MySQL portion of the site. You will have to do a little digging into the files. But don't worry! I saved you a backup here on github! <3

First things first, if you decide to keep the post request feature, and implement your own version of post request support, where I address line deletion is probably where you need to do your modifications. 

Getting into the deletions though, we need to open the `portfolio.go` file. You should see `http.HandleFunc("/formsubmissionhandler", addWebForm)` around line 18. You can go ahead and delete that. Scroll down about halfway to the bottom around line 98, you should see something like `type Form struct {`, and you can delete this line and everything below it. And finally, we must go to the top and delete all the imports that were used. In this case, it should see `_ "github.com/go-sql-driver/mysql"` around line 12, this can be deleted as well. That's it for disabling the MySQL features.

### Without TLS With MySQL

This is probably the easiest way to set up the website. You do not need to set up any certificates, or delete any code. But you do, need to change one line of code to submit your forms properly. Locate the file at `/lib/JS/submit.js` and change the line url around line 22, to `http://your.domain/formsubmissionhandler` which will save any forms submitted using the server. (Please notice the absence of 's' in http. As https is port 443, and http is port 80)

## Starting The Server

Please refer back to [Usage](https://github.com/WesternPine/Portfolio#startup-requires-golang-optional-mysql) on how to start the server.


## Usage

Go ahead and visit your site! If the site is hosted, please use the ip or domain. If the site is on your computer, go to `localhost` in your browser. Please also note that is you use TLS, you must put `https://` before your ip/domain.

## Final Notes

And that's it! There's nothin to it if you somewhat understand what you're doing! There are a lot of technical aspects of this project that some may find complicated. If you do not understand the basic concepts of golang or any web development, I highly suggest going to do a little research before diving into this project. If you have any other questions, comments, or concerns, feel free to contact me here on github or use my website in my profile. Thank you!

License
----

[MIT](https://choosealicense.com/)
