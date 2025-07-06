// API Configuration
const API_BASE_URL = 'http://localhost:8080/api';

// State management
let currentUser = null;
let authToken = null;

// Initialize the application
document.addEventListener('DOMContentLoaded', function() {
    // Check for existing authentication
    checkAuthStatus();
    
    // Load initial data
    loadDestinations();
    loadServices();
    
    // Set up event listeners
    setupEventListeners();
});

// Authentication functions
function checkAuthStatus() {
    authToken = localStorage.getItem('authToken');
    if (authToken) {
        // TODO: Validate token with backend
        updateAuthUI(true);
    }
}

function updateAuthUI(isAuthenticated) {
    const authButtons = document.querySelector('.nav-auth');
    if (isAuthenticated) {
        authButtons.innerHTML = `
            <button class="btn btn-outline" onclick="logout()">Logout</button>
            <button class="btn btn-primary" onclick="showBookings()">My Bookings</button>
        `;
    } else {
        authButtons.innerHTML = `
            <button class="btn btn-outline" onclick="openModal('loginModal')">Sign In</button>
            <button class="btn btn-primary" onclick="openModal('registerModal')">Sign Up</button>
        `;
    }
}

function logout() {
    localStorage.removeItem('authToken');
    authToken = null;
    currentUser = null;
    updateAuthUI(false);
}

// API functions
async function apiRequest(endpoint, options = {}) {
    const url = `${API_BASE_URL}${endpoint}`;
    const config = {
        headers: {
            'Content-Type': 'application/json',
            ...options.headers
        },
        ...options
    };
    
    if (authToken) {
        config.headers.Authorization = `Bearer ${authToken}`;
    }
    
    try {
        const response = await fetch(url, config);
        const data = await response.json();
        
        if (!response.ok) {
            throw new Error(data.message || 'API request failed');
        }
        
        return data;
    } catch (error) {
        console.error('API Error:', error);
        throw error;
    }
}

// Load destinations
async function loadDestinations() {
    try {
        const response = await apiRequest('/destinations');
        const destinations = response.data;
        renderDestinations(destinations);
    } catch (error) {
        console.error('Error loading destinations:', error);
        document.getElementById('destinationsGrid').innerHTML = 
            '<div class="error">Failed to load destinations. Please try again later.</div>';
    }
}

// Render destinations
function renderDestinations(destinations) {
    const grid = document.getElementById('destinationsGrid');
    grid.innerHTML = destinations.map(destination => `
        <div class="destination-card">
            <img src="${destination.image_url}" alt="${destination.name}" class="destination-image">
            <div class="destination-info">
                <div class="destination-header">
                    <h3 class="destination-name">${destination.name}</h3>
                    <div class="destination-rating">
                        <i class="fas fa-star"></i>
                        <span>${destination.rating}</span>
                    </div>
                </div>
                <p class="destination-description">${destination.description}</p>
                <div class="destination-deals">${destination.deals_count} deals</div>
            </div>
        </div>
    `).join('');
}

// Load services
async function loadServices() {
    try {
        const response = await apiRequest('/services');
        const services = response.data;
        renderServices(services);
    } catch (error) {
        console.error('Error loading services:', error);
        document.getElementById('servicesGrid').innerHTML = 
            '<div class="error">Failed to load services. Please try again later.</div>';
    }
}

// Render services
function renderServices(services) {
    const grid = document.getElementById('servicesGrid');
    grid.innerHTML = services.map(service => `
        <div class="service-card">
            <img src="${service.image_url}" alt="${service.name}" class="service-icon">
            <h3 class="service-name">${service.name}</h3>
            <p class="service-description">${service.description}</p>
        </div>
    `).join('');
}

// Modal functions
function openModal(modalId) {
    document.getElementById(modalId).style.display = 'block';
}

function closeModal(modalId) {
    document.getElementById(modalId).style.display = 'none';
}

function switchModal(currentModalId, targetModalId) {
    closeModal(currentModalId);
    openModal(targetModalId);
}

// Event listeners
function setupEventListeners() {
    // Login form
    document.getElementById('loginForm').addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const email = document.getElementById('loginEmail').value;
        const password = document.getElementById('loginPassword').value;
        
        try {
            const response = await apiRequest('/auth/login', {
                method: 'POST',
                body: JSON.stringify({ email, password })
            });
            
            authToken = response.data.token;
            currentUser = response.data.user;
            localStorage.setItem('authToken', authToken);
            
            updateAuthUI(true);
            closeModal('loginModal');
            showNotification('Login successful!', 'success');
        } catch (error) {
            showNotification(error.message, 'error');
        }
    });
    
    // Register form
    document.getElementById('registerForm').addEventListener('submit', async function(e) {
        e.preventDefault();
        
        const firstName = document.getElementById('registerFirstName').value;
        const lastName = document.getElementById('registerLastName').value;
        const email = document.getElementById('registerEmail').value;
        const phone = document.getElementById('registerPhone').value;
        const password = document.getElementById('registerPassword').value;
        
        try {
            const response = await apiRequest('/auth/register', {
                method: 'POST',
                body: JSON.stringify({
                    first_name: firstName,
                    last_name: lastName,
                    email,
                    phone,
                    password
                })
            });
            
            authToken = response.data.token;
            currentUser = response.data.user;
            localStorage.setItem('authToken', authToken);
            
            updateAuthUI(true);
            closeModal('registerModal');
            showNotification('Registration successful!', 'success');
        } catch (error) {
            showNotification(error.message, 'error');
        }
    });
    
    // Booking forms
    document.getElementById('hotelForm').addEventListener('submit', function(e) {
        e.preventDefault();
        if (!authToken) {
            showNotification('Please login to make a booking', 'error');
            openModal('loginModal');
            return;
        }
        // TODO: Implement hotel booking
        showNotification('Hotel booking feature coming soon!', 'info');
    });
    
    document.getElementById('busForm').addEventListener('submit', function(e) {
        e.preventDefault();
        if (!authToken) {
            showNotification('Please login to make a booking', 'error');
            openModal('loginModal');
            return;
        }
        // TODO: Implement bus booking
        showNotification('Bus booking feature coming soon!', 'info');
    });
    
    document.getElementById('visaForm').addEventListener('submit', function(e) {
        e.preventDefault();
        if (!authToken) {
            showNotification('Please login to request assistance', 'error');
            openModal('loginModal');
            return;
        }
        // TODO: Implement visa assistance
        showNotification('Visa assistance feature coming soon!', 'info');
    });
    
    // Close modals when clicking outside
    window.addEventListener('click', function(e) {
        if (e.target.classList.contains('modal')) {
            e.target.style.display = 'none';
        }
    });
    
    // Mobile menu toggle
    document.querySelector('.nav-toggle').addEventListener('click', function() {
        const navMenu = document.querySelector('.nav-menu');
        navMenu.style.display = navMenu.style.display === 'flex' ? 'none' : 'flex';
    });
}

// Show bookings
async function showBookings() {
    if (!authToken) {
        showNotification('Please login to view bookings', 'error');
        openModal('loginModal');
        return;
    }
    
    try {
        const response = await apiRequest('/bookings');
        const bookings = response.data;
        
        if (bookings.length === 0) {
            showNotification('No bookings found', 'info');
            return;
        }
        
        // TODO: Create a bookings modal or page
        console.log('User bookings:', bookings);
        showNotification(`You have ${bookings.length} booking(s)`, 'info');
    } catch (error) {
        showNotification('Failed to load bookings', 'error');
    }
}

// Notification system
function showNotification(message, type = 'info') {
    // Remove existing notification
    const existingNotification = document.querySelector('.notification');
    if (existingNotification) {
        existingNotification.remove();
    }
    
    // Create notification element
    const notification = document.createElement('div');
    notification.className = `notification notification-${type}`;
    notification.innerHTML = `
        <span>${message}</span>
        <button onclick="this.parentElement.remove()">&times;</button>
    `;
    
    // Add notification styles
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        right: 20px;
        background: ${type === 'error' ? '#ef4444' : type === 'success' ? '#10b981' : '#3b82f6'};
        color: white;
        padding: 1rem 1.5rem;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        z-index: 3000;
        display: flex;
        align-items: center;
        gap: 1rem;
        animation: slideIn 0.3s ease-out;
    `;
    
    // Add animation keyframes
    if (!document.querySelector('#notification-styles')) {
        const style = document.createElement('style');
        style.id = 'notification-styles';
        style.textContent = `
            @keyframes slideIn {
                from {
                    transform: translateX(100%);
                    opacity: 0;
                }
                to {
                    transform: translateX(0);
                    opacity: 1;
                }
            }
        `;
        document.head.appendChild(style);
    }
    
    document.body.appendChild(notification);
    
    // Auto-remove after 5 seconds
    setTimeout(() => {
        if (notification.parentElement) {
            notification.remove();
        }
    }, 5000);
}

// Smooth scrolling for navigation links
document.addEventListener('DOMContentLoaded', function() {
    const navLinks = document.querySelectorAll('.nav-link');
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const targetId = this.getAttribute('href').substring(1);
            const targetElement = document.getElementById(targetId);
            if (targetElement) {
                targetElement.scrollIntoView({ behavior: 'smooth' });
            }
        });
    });
});

// Add loading states
function showLoading(elementId) {
    const element = document.getElementById(elementId);
    if (element) {
        element.innerHTML = '<div class="loading">Loading...</div>';
    }
}

function hideLoading(elementId) {
    const element = document.getElementById(elementId);
    if (element) {
        const loadingElement = element.querySelector('.loading');
        if (loadingElement) {
            loadingElement.remove();
        }
    }
}

// Handle form validation
function validateForm(formData) {
    const errors = [];
    
    if (!formData.email || !formData.email.includes('@')) {
        errors.push('Please enter a valid email address');
    }
    
    if (!formData.password || formData.password.length < 6) {
        errors.push('Password must be at least 6 characters long');
    }
    
    return errors;
}

// Initialize page
console.log('Nomado frontend initialized');
