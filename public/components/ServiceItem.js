export class ServiceItemComponent extends HTMLElement {
    constructor(service) {
        super();
        this.service = service;
    }
    connectedCallback() {
        this.innerHTML = `
            <a href="#">
                <article>
                    <img src="${this.service.image}" alt="${this.service.name}">
                    <p>${this.service.name}</p>
                </article>
            </a>
        `;
    }
}

customElements.define("service-item", ServiceItemComponent);