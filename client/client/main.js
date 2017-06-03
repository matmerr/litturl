Vue.use(VueMaterial)

Vue.material.registerTheme({
  default: {
    primary: {
      color: 'light-blue',
      hue: 700
    },
    accent: 'red'
  }
})



var App = new Vue({
  el: '#app',
  data: () => ({
    vertical: 'bottom',
    horizontal: 'center',
    duration: 4000,
    urli: "",
    message: ""
  }),
  methods: {
    open() {

      this.message = Post(this.urli);
      this.$refs.snackbar.open();
    }
  }
})

function Post(newurl) {


  var xhr = new XMLHttpRequest();
  xhr.open("POST", "http://10.0.1.250:8000/add", true);
  xhr.setRequestHeader('Content-Type', 'application/json');

  xhr.send(JSON.stringify({
    apikey: "ok",
    url: "okok"
  }));
  return ;
}


function F() {
  //document.getElementById("demo").innerHTML = "Hello World";
  var xmlhttp = new XMLHttpRequest();
  var url = "http://192.168.91.137:8000/StudKentFoo/json";

  xmlhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {}
  };
  xmlhttp.open("GET", url, false);
  xmlhttp.send();
  return xmlhttp.responseText;
}
