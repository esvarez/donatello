'use strict';

let taskContainer = document.getElementById('task-container');

chrome.storage.sync.get('tasks', function(data) {
  let tasks = data.tasks
  tasks.sort(compare)

  let fragment = document.createDocumentFragment()
  tasks.forEach(task => {
    let label = document.createElement('label'),
        input = document.createElement('input'),
        spanCheck = document.createElement('span'),
        spanText = document.createElement('span');

    input.type = 'checkbox'
    if(task.done){
      input.checked='checked'
    }
    spanCheck.classList.add('checkmark')
    spanText.innerText = task.name;  
    label.classList.add('container')      
    label.appendChild(input)
    label.appendChild(spanCheck)
    label.appendChild(spanText)
    fragment.appendChild(label)
  });
  taskContainer.appendChild(fragment)
})

function compare(a, b) {
  return (a.done)? 1 : -1;
}
/*

let changeColor = document.getElementById('changeColor');

changeColor.onclick = function(element) {
  let color = element.target.value;
  chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
    chrome.tabs.executeScript(
        tabs[0].id,
        {code: 'document.body.style.backgroundColor = "' + color + '";'});
  });
};

chrome.storage.sync.get('color', function(data) {
  changeColor.style.backgroundColor = data.color;
  changeColor.setAttribute('value', data.color);
});
*/
