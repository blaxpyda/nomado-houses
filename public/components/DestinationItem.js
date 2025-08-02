export class DestinationItemComponent extends HTMLElement {
    constructor(destination) {
        super();
        this.destination = destination;
    }
    connectedCallback() {
        this.innerHTML = `
            <a href="#">
                <article>
                    <img src="${this.destination.image}" alt="${this.destination.name}">
                    <p>${this.destination.name}</p>
                </article>
            </a>
        `;
    }
}
customElements.define('destination-item', DestinationItemComponent);