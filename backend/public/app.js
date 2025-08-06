import { API } from "./services/API.js";
import { HomePage } from "./components/HomePage.js";

window.addEventListener("DOMContentLoaded", event => {
    document.querySelector("main").appendChild(new HomePage());
    
    // Mobile menu functionality
    const mobileMenuButton = document.getElementById('mobile-menu-button');
    const mobileMenu = document.getElementById('mobile-menu');
    
    if (mobileMenuButton && mobileMenu) {
        mobileMenuButton.addEventListener('click', () => {
            mobileMenu.classList.toggle('hidden');
        });
    }
    
    // Smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function (e) {
            e.preventDefault();
            const target = document.querySelector(this.getAttribute('href'));
            if (target) {
                target.scrollIntoView({
                    behavior: 'smooth',
                    block: 'start'
                });
            }
        });
    });
    
    // Add scroll effect to header
    const header = document.querySelector('header');
    if (header) {
        window.addEventListener('scroll', () => {
            if (window.scrollY > 100) {
                header.classList.add('bg-white/95');
                header.classList.remove('bg-white/90');
            } else {
                header.classList.add('bg-white/90');
                header.classList.remove('bg-white/95');
            }
        });
    }
});

window.app = {
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
        // TODO: Implement search functionality
    },
    api: API
}