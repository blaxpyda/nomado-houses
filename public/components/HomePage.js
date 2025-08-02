import { ServiceTypeItemComponent } from "./ServiceTypeItem.js";
import { DestinationItemComponent } from "./DestinationItem.js";
import { API } from "../services/API.js";

export class HomePage extends HTMLElement { // <home-page>

    async render() {
        const allServices = await API.getAllServices();
        renderServicesInList(allServices, document.querySelector("#service-types ul"));

        const allDestinations = await API.getAllDestinations();
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