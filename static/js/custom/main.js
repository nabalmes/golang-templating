var signUpBtn = document.querySelector('.signup-btn');
var CloseBtn = document.querySelector('.close-modal');
var modal = document.querySelector('.modal-container');
var signinSection = document.querySelector('.signin-section')
signUpBtn.addEventListener('click', function(){
    modal.classList.add('show')
    signinSection.display = 'none';
})
CloseBtn.addEventListener('click', function(){
    modal.classList.remove('show')
    signinSection.display = 'flex';
})