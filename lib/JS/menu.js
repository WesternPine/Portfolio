const navSlide = () => {
  const menu = document.querySelector('.menu');
  const nav = document.querySelector('.nav-links');
  const navLinks = document.querySelectorAll('.nav-links li');
  
  
  menu.addEventListener('click', () => {
    
    //navigation animation
    nav.classList.toggle('menu-active');
    
    //animated links
    navLinks.forEach((link, index) => {
      if(link.style.animation){
        link.style.animation = '';
      } else {
        link.style.animation = `navLinkFade 0.5s ease-in-out forwards ${index * 200 + 500}ms`;
      }
    });
    
    //menu animation
    menu.classList.toggle('toggle');
    
  });
  
  
}

navSlide();