import { ServiceTypeItemComponent } from "./ServiceTypeItem.js";
import { DestinationItemComponent } from "./DestinationItem.js";
import { API } from "../services/API.js";

export class HomePage extends HTMLElement { // <home-page>

    async render() {
        const servicesResponse = await API.getAllServices();
        console.log('Services response:', servicesResponse); // Debug log
        const allServices = servicesResponse.data || servicesResponse; // Extract the data array
        console.log('Extracted services:', allServices); // Debug log
        renderServicesInList(allServices, document.querySelector("#service-types ul"));

        const destinationsResponse = await API.getAllDestinations();
        console.log('Destinations response:', destinationsResponse); // Debug log
        const allDestinations = destinationsResponse.data || destinationsResponse; // Extract the data array
        console.log('Extracted destinations:', allDestinations); // Debug log
        renderDestinationsInList(allDestinations, document.querySelector("#destinations ul"));

        function renderServicesInList(services, ul) {
            ul.innerHTML = '';
            services.forEach(service => {
                const li = document.createElement("li");
                li.appendChild(new ServiceTypeItemComponent(service));
                ul.appendChild(li);
            });
        }
        function renderDestinationsInList(destinations, ul) {
            ul.innerHTML = '';
            destinations.forEach(destination => {
                const li = document.createElement("li");
                li.appendChild(new DestinationItemComponent(destination));
                ul.appendChild(li);
            });
        }
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