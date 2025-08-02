export class ServiceTypeItemComponent extends HTMLElement {
    constructor(serviceType) {
        super();
        this.serviceType = serviceType;
    }
    connectedCallback() {
        this.innerHTML = `
            <a href="#">
                <article>
                    <img src="${this.serviceType.image}" alt="${this.serviceType.name}">
                    <p>${this.serviceType.name}</p>
                </article>
            </a>
        `;
    }
}

customElements.define("service-type-item", ServiceTypeItemComponent);