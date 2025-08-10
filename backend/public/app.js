import { API } from "./services/API.js";
import { HomePage } from "./components/HomePage.js";

// Authentication utility functions
const auth = {
    // Get stored token
    getToken() {
        return localStorage.getItem('authToken');
    },

    // Get stored user data
    getUser() {
        const userData = localStorage.getItem('userData');
        return userData ? JSON.parse(userData) : null;
    },

    // Check if user is logged in
    isLoggedIn() {
        return !!this.getToken();
    },

    // Logout user
    logout() {
        localStorage.removeItem('authToken');
        localStorage.removeItem('userData');
        window.location.href = '/';
    },

    // Update UI based on authentication state
    updateUI() {
        const loginButton = document.querySelector('button[onclick="location.href=\'/login\'"]');
        const registerButton = document.querySelector('button[onclick="location.href=\'/register\'"]');
        
        if (this.isLoggedIn()) {
            const user = this.getUser();
            if (loginButton && registerButton) {
                // Replace login/register buttons with user menu
                const userMenu = document.createElement('div');
                userMenu.className = 'flex items-center space-x-3';
                userMenu.innerHTML = `
                    <div class="hidden md:flex items-center space-x-2">
                        <span class="text-gray-700">Welcome, ${user.first_name}!</span>
                    </div>
                    <div class="relative group">
                        <button class="flex items-center space-x-2 px-4 py-2 bg-gray-100 hover:bg-gray-200 rounded-full transition-colors duration-200">
                            <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-purple-600 rounded-full flex items-center justify-center text-white font-semibold text-sm">
                                ${user.first_name.charAt(0).toUpperCase()}${user.last_name.charAt(0).toUpperCase()}
                            </div>
                            <svg class="w-4 h-4 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                            </svg>
                        </button>
                        <div class="absolute right-0 mt-2 w-48 bg-white rounded-lg shadow-lg border border-gray-200 opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all duration-200 z-50">
                            <div class="p-3 border-b border-gray-100">
                                <p class="font-medium text-gray-900">${user.first_name} ${user.last_name}</p>
                                <p class="text-sm text-gray-600">${user.email}</p>
                            </div>
                            <div class="py-1">
                                <a href="/account" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50">My Account</a>
                                <a href="/bookings" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-50">My Bookings</a>
                                <button onclick="window.app.auth.logout()" class="block w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50">Logout</button>
                            </div>
                        </div>
                    </div>
                `;
                
                // Replace the auth buttons with user menu
                const authContainer = loginButton.parentElement;
                authContainer.innerHTML = '';
                authContainer.appendChild(userMenu);
            }
        }
    }
};

window.addEventListener("DOMContentLoaded", event => {
    document.querySelector("main").appendChild(new HomePage());
    
    // Update UI based on authentication state
    auth.updateUI();
    
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
    auth: auth,
    api: API
}