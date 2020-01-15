let taskContainer = document.getElementById('task-container');

const search = document.getElementById('search-button')

search.onclick = function() {
  search.parentElement.classList.toggle('open')
}

/*
function setTask() {
  //let tasks = data.tasks
  let tasks = [
   {id: 1, name: 'One', done: true},
   {id: 2, name: 'Two Two', done: false},
   {id: 3, name: 'Three Three Three', done: true},
   {id: 4, name: 'four four four four', done: true},
 ]
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
}

function compare(a, b) {
  return (a.done)? 1 : -1;
}


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

/* button.addEventListener("click", function(){ 
   console.log(checkbox)
}); 

setTask()
*/