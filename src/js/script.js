const taskContainer = document.getElementById('task-container'),
      search = document.getElementById('add-button'),
      inputTask = document.getElementById('input-task')

const dodos = []

search.onclick = function() {
  search.parentElement.classList.toggle('open')
  if (inputTask.value) {
    console.log('no vacio')  
    let label = document.createElement('label'),
        input = document.createElement('input'),
        spanCheck = document.createElement('span'),
        spanText = document.createElement('span');

    input.type = 'checkbox'    
    spanCheck.classList.add('checkmark')
    spanText.innerText = inputTask.value  
    label.classList.add('container')      
    label.appendChild(input)
    label.appendChild(spanCheck)
    label.appendChild(spanText)    

    insertAfter(label, dodos[dodos.length - 1])

    dodos.push(label) 
    inputTask.value = ''       
  } else {
    console.log('vacio')  
  }
  console.log(dodos)
  //addTask(inputTask.value)  
}

function insertAfter(el, referenceNode) {
    referenceNode.parentNode.insertBefore(el, referenceNode.nextSibling);
}

let tasks = [
  {id: 1, name: 'One', done: true},
  {id: 2, name: 'Two Two', done: false},
  {id: 3, name: 'Three Three Three', done: true},
  {id: 4, name: 'four four four four', done: true},
]

function addTask(task) {
   tasks.push(
    {id: 1, name: task, done: false},
   )
}

function setTask() {
  //let tasks = data.tasks  
  tasks.sort(compare)

  let fragment = document.createDocumentFragment()
  tasks.forEach(task => {
    let label = document.createElement('label'),
        input = document.createElement('input'),
        spanCheck = document.createElement('span'),
        spanText = document.createElement('span');

    input.type = 'checkbox'    
    spanCheck.classList.add('checkmark')
    spanText.innerText = task.name;  
    label.classList.add('container')      
    label.appendChild(input)
    label.appendChild(spanCheck)
    label.appendChild(spanText)
    fragment.appendChild(label)
    if (task.done){
      input.checked='checked'      
    } else {
      dodos.push(label)
    }
  });
  taskContainer.appendChild(fragment)
}

function createTask(task) {
  
}

function compare(a, b) {
  return (a.done)? 1 : -1;
}

setTask()


/*
console.log('...')
const checkbox = document.getElementById("checkbox"),
btn = document.getElementById("btn"),
form = document.getElementById("form"),
btm = document.getElementById("btm"),
label = document.getElementById("label")

btn.addEventListener("click", function(){       
      //checkbox.checked = true
      //checkbox.removeAttribute('checked')
      //form.submit()
   console.log('true')             
   checkbox.checked = true
});

btm.addEventListener("click", function(){ 
   console.log('false')       
   checkbox.checked = false           
});

label.addEventListener('click', function() {
   console.log('click')
})

 button.addEventListener("click", function(){ 
   console.log(checkbox)
}); 


*/