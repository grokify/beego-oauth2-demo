# Beego OAuth 2.0 Example

[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

A basic [Beego](https://beego.me/) example app for OAuth 2.0 using Google and Facebook. 

The following are used in this example app:

* OAuth 2.0
  * [OAuth2](https://github.com/golang/oauth2)
  * [OAuth2 Util](https://github.com/grokify/oauth2-util-go)
* Rendering
  * [Bootstrap](http://getbootstrap.com/)
  * [Font Awesome](http://fontawesome.io/)
  * [Social Buttons](https://lipis.github.io/bootstrap-social/)
  * [Quicktemplate](https://github.com/valyala/quicktemplate)

## Usage

### App Hostname / OAuth 2.0 Redirect URI

Decide on the hostname for your app which will be used in the OAuth 2.0 URI.

The redirect URI is `oauth2callback`. For example, if your app URL is `https://example.com`, your redirect URI will be `https://example.com/oauth2callback`.

If you do not have an Internet accessible hostname, you can use [ngrok](https://ngrok.com/) for testing.

### Configure Google and Facebook

Login to the Google and Facebbook developer consoles to configure your apps with your OAuth 2.0 rediect URIs. During this process, you will get your OAuth 2.0 client ID and client secrets for Google and Facebook.

### Install & Configure App

1. Clone the repo
2. Install the dependencies
3. Set up your `app.conf` file.
4. Run the app

You can do the above with the following steps. See the next section for more detailed information on the configuration file.

```bash
$ git clone https://github.com/grokify/beego-oauth2-demo
$ go get ./...
$ cd beego-oauth2-demo/conf
$ cp app.conf.sample app.conf
$ vim app.conf
$ bee run
```

#### Configuration File

This is a basic Beego config file with additional properties. Specifically, it has the following 3 properties:

| Property | Description |
|----------|-------------|
| `oauth2configgoogle` | The Google Developers Console `client_credentials.json` file information. This is read directly by OAuth2's `google.ConfigFromJSON()` function. |
| `oauth2configfacebook` | This uses the format used by `oauth2.Config` |
| `oauth2redirecturi` | This app's OAuth2 redirect URI which has the path `/oauth2callback`, e.g. `https://example.com/oauth2callback` |

 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/beego-oauth2-demo
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/beego-oauth2-demo
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/beego-oauth2-demo
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/beego-oauth2-demo/blob/master/LICENSE.md
