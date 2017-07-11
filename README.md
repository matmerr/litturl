# littURL
[![Docker Pulls](https://img.shields.io/docker/pulls/matmerr/litturl.svg)][![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fmatmerr%2Flitturl.svg?type=shield)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fmatmerr%2Flitturl?ref=badge_shield)
() [![](https://images.microbadger.com/badges/image/matmerr/litturl.svg)](https://microbadger.com/images/matmerr/litturl "Get your own image badge on microbadger.com")
> A little URL shortener.
## Features:
- Self Hosted URL Shortener
- Simple installation with Docker
- Material Design using [VueMaterial](http://vuematerial.io/#/) and [Vue.JS](https://vuejs.org/)


## Get up and running with Docker Compose
> If you don't have Docker Compose installed, [check this out](https://docs.docker.com/compose/install/#install-as-a-container)

```
# clone the repository
git clone https://github.com/matmerr/litturl

# cd to repo
cd litturl

# bring it up
docker-compose up -d
```
## Or if you have a Redis instance
> note: by default uses db 0
```
# clone the repository
git clone https://github.com/matmerr/litturl

# cd to repo
cd litturl

# bring it up with or with persistant storage
docker run -d -p 8001:8001 -v /host/dir/conf_dir:/go/src/github.com/matmerr/litturl/conf matmerr/litturl
```


## Screenshots

### **Inital Setup:**
[![settings](docs/images/initial_setup.png)]()
### **Home:**
[![home](docs/images/home.png)]()
### **Settings:**
[![settings](docs/images/settings.png)]()

## TODO's
- Multiuser authentication
- BYODB for others besides Redis
- Google Analytics / Click Statistics


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fmatmerr%2Flitturl.svg?type=large)](https://app.fossa.io/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fmatmerr%2Flitturl?ref=badge_large)