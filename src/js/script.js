console.log('...')
const checkbox = document.getElementById("checkbox"),
btn = document.getElementById("btn"),
form = document.getElementById("form"),
btm = document.getElementById("btm")

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
}); */