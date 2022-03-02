var xhr = new XMLHttpRequest()
var url = 'http://localhost:8888/'

const handler = () => {
  // コンソールに出力
  console.log(xhr.responseText)
}

const postJsonRequest = () => {
  xhr.open('POST', url)
  xhr.onloadend = handler
  xhr.setRequestHeader('Content-Type', 'application/json')
  xhr.send()
}

document.addEventListener('DOMContentLoaded', () => {
  postJsonRequest()
})
