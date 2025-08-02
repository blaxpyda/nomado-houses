import { ServiceItemComponent } from "./ServiceItem";

export class HomePage extends HTMLElement { // <home-page>

    async render() {
        const allServices = await API.getAllServices();
        renderServicesInList(allServices, document.querySelector("#top-10 ul"));

        const allDestinations = await API.getAllDestinations();
        renderServicesInList(allDestinations, document.querySelector("#top-destinations ul"));

        function renderServicesInList(services, ul) {
            ul.innerHTML = '';
            services.forEach(service => {
                const li = document.createElement("li");
                li.appendChild(new ServiceItemComponent(service));
                ul.appendChild(li);
            });
        }
    }
    connectedCallback() {
        const template = document.getElementById("template-home");
        const content = template.contentEditable.cloneNode(true);
        this.appendChild(template);

    }
}
customElements.define('home-page', HomePage);