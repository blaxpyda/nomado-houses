import { ServiceTypeItemComponent } from "./ServiceTypeItem.js";
import { DestinationItemComponent } from "./DestinationItem.js";
import { API } from "../services/API.js";

export class HomePage extends HTMLElement { // <home-page>

    async render() {
        const servicesContainer = document.querySelector("#service-types ul");
        const destinationsContainer = document.querySelector("#destinations ul");
        
        // Show loading state
        if (servicesContainer) {
            servicesContainer.innerHTML = this.getLoadingHTML();
        }
        if (destinationsContainer) {
            destinationsContainer.innerHTML = this.getLoadingHTML();
        }

        try {
            // Fetch services
            const servicesResponse = await API.getAllServices();
            console.log('Services response:', servicesResponse);
            const allServices = servicesResponse.data || servicesResponse;
            console.log('Extracted services:', allServices);
            if (servicesContainer) {
                renderServicesInList(allServices, servicesContainer);
            }

            // Fetch destinations
            const destinationsResponse = await API.getAllDestinations();
            console.log('Destinations response:', destinationsResponse);
            const allDestinations = destinationsResponse.data || destinationsResponse;
            console.log('Extracted destinations:', allDestinations);
            if (destinationsContainer) {
                renderDestinationsInList(allDestinations, destinationsContainer);
            }
        } catch (error) {
            console.error('Error fetching data:', error);
            if (servicesContainer) {
                servicesContainer.innerHTML = this.getErrorHTML('Failed to load services');
            }
            if (destinationsContainer) {
                destinationsContainer.innerHTML = this.getErrorHTML('Failed to load destinations');
            }
        }

        function renderServicesInList(services, ul) {
            ul.innerHTML = '';
            services.forEach(service => {
                ul.appendChild(new ServiceTypeItemComponent(service));
            });
        }
        function renderDestinationsInList(destinations, ul) {
            ul.innerHTML = '';
            destinations.forEach(destination => {
                ul.appendChild(new DestinationItemComponent(destination));
            });
        }
    }

    getLoadingHTML() {
        return `
            <div class="col-span-full flex justify-center items-center py-12">
                <div class="flex flex-col items-center space-y-4">
                    <div class="animate-spin rounded-full h-12 w-12 border-4 border-blue-500 border-t-transparent"></div>
                    <p class="text-gray-600 font-medium">Loading amazing destinations...</p>
                </div>
            </div>
        `;
    }

    getErrorHTML(message) {
        return `
            <div class="col-span-full flex justify-center items-center py-12">
                <div class="text-center">
                    <div class="w-16 h-16 mx-auto mb-4 bg-red-100 rounded-full flex items-center justify-center">
                        <svg class="w-8 h-8 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                    </div>
                    <p class="text-gray-600 font-medium">${message}</p>
                    <button onclick="location.reload()" class="mt-4 px-6 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors">
                        Try Again
                    </button>
                </div>
            </div>
        `;
    }
    connectedCallback() {
        const template = document.getElementById("template-home");
        const content = template.content.cloneNode(true);
        this.appendChild(content);
        
        // Call render after the template is added to the DOM
        this.render();
    }
}
customElements.define('home-page', HomePage);