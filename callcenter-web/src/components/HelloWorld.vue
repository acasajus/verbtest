<template>
  <div>
    <div class='req'>
      <button @click="doUnauthorized">Unauthorized request</button>
    </div> 
    <div class='req'>
      <button @click="doDirectNLU">Send text to nlu</button>
      <input type="text" v-model="text_nlu">
    </div> 
    <div class='req'>
      <button @click="doDM">Send text to dialog manager</button>
      <input type="text" v-model="text_dm"> Forward to
      <span v-for="svc in ['nlu', 'task', 'direct', 'knowledge']">
        <input type="radio" :id="svc" :value="svc" v-model="svc_dm">
        <label :for="svc">{{svc}}</label>
      </span>
    </div> 
    <p>
      {{response}}
    </p>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'HelloWorld',
  data() {
    return {
      response: 'no response',
      text_nlu: 'text to send',
      text_dm: 'more input',
      svc_dm: 'task'
    }
  },
  methods: {
    doUnauthorized(e) {
      e.preventDefault()
      this.response = 'making request....'
      axios.get('http://localhost:8000/nlu')
        .then(resp => (this.response = 'ok ' + resp))
        .catch(err => this.response = 'ko ' + JSON.stringify(err))
    },
    doDirectNLU(e) {
      e.preventDefault()
      this.response = 'making request with text ' + this.text_nlu
      axios.post('http://localhost:8000/nlu', {text: this.text_nlu}, {headers: {'Authorization': 'Bearer aaa'}})
        .then(resp => (this.response = 'OK ' + resp.data.data))
        .catch(err => this.response = 'KO ' + JSON.stringify(err))
    },
    doDM(e) {
      e.preventDefault()
      this.response = 'making request with text ' + this.text_dm
      axios.post('http://localhost:8000/dialog', {text:this.svc_dm + ': ' + this.text_dm}, {headers: {'Authorization': 'Bearer aaa'}})
        .then(resp => (this.response = 'OK ' + resp.data.data))
        .catch(err => this.response = 'KO ' + JSON.stringify(err))
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div {
  margin: 10px;
}
input {
  margin: 10px;
}
</style>
