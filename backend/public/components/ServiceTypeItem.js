export class ServiceTypeItemComponent extends HTMLElement {
    constructor(serviceType) {
        super();
        this.serviceType = serviceType;
    }

    getServiceIcon(serviceName) {
        const name = serviceName.toLowerCase();
        
        if (name.includes('hotel') || name.includes('guesthouse')) {
            return `
                <svg class="w-full h-full text-blue-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M19 7h-3V6a4 4 0 0 0-8 0v1H5a1 1 0 0 0-1 1v11a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-1-1zM10 6a2 2 0 0 1 4 0v1h-4V6zm8 13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9h2v1a1 1 0 0 0 2 0V9h4v1a1 1 0 0 0 2 0V9h2v10z"/>
                    <path d="M9 12h6v2H9zm0 3h6v2H9z"/>
                </svg>`;
        } else if (name.includes('visa')) {
            return `
                <svg class="w-full h-full text-purple-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V5h14v14z"/>
                    <path d="M7 7h10v2H7zm0 3h10v2H7zm0 3h7v2H7z"/>
                </svg>`;
        } else if (name.includes('flight')) {
            return `
                <svg class="w-full h-full text-sky-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M21 16v-2l-8-5V3.5c0-.83-.67-1.5-1.5-1.5S10 2.67 10 3.5V9l-8 5v2l8-2.5V19l-2 1.5V22l3.5-1 3.5 1v-1.5L13 19v-5.5l8 2.5z"/>
                </svg>`;
        } else if (name.includes('bus')) {
            return `
                <svg class="w-full h-full text-green-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M4 16c0 .88.39 1.67 1 2.22V20c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-1h8v1c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-1.78c.61-.55 1-1.34 1-2.22V6c0-3.5-3.58-4-8-4s-8 .5-8 4v10zm3.5 1c-.83 0-1.5-.67-1.5-1.5S6.67 14 7.5 14s1.5.67 1.5 1.5S8.33 17 7.5 17zm9 0c-.83 0-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5 1.5.67 1.5 1.5-.67 1.5-1.5 1.5zm1.5-6H6V6h12v5z"/>
                </svg>`;
        } else if (name.includes('car') || name.includes('ride')) {
            return `
                <svg class="w-full h-full text-red-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M18.92 6.01C18.72 5.42 18.16 5 17.5 5h-11c-.66 0-1.21.42-1.42 1.01L3 12v8c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-1h12v1c0 .55.45 1 1 1h1c.55 0 1-.45 1-1v-8l-2.08-5.99zM6.5 16c-.83 0-1.5-.67-1.5-1.5S5.67 13 6.5 13s1.5.67 1.5 1.5S7.33 16 6.5 16zm11 0c-.83 0-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5 1.5.67 1.5 1.5-.67 1.5-1.5 1.5zM5 11l1.5-4.5h11L19 11H5z"/>
                </svg>`;
        } else if (name.includes('love')) {
            return `
                <svg class="w-full h-full text-pink-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"/>
                </svg>`;
        } else if (name.includes('nomad') || name.includes('little')) {
            return `
                <svg class="w-full h-full text-orange-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 2c1.1 0 2 .9 2 2s-.9 2-2 2-2-.9-2-2 .9-2 2-2zm9 7h-6v13h-2v-6h-2v6H9V9H3V7h18v2z"/>
                </svg>`;
        } else if (name.includes('event') || name.includes('retreat')) {
            return `
                <svg class="w-full h-full text-indigo-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M17 12h-5v5h5v-5zM16 1v2H8V1H6v2H5c-1.11 0-1.99.9-1.99 2L3 19c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2h-1V1h-2zm3 18H5V8h14v11z"/>
                </svg>`;
        } else if (name.includes('job')) {
            return `
                <svg class="w-full h-full text-slate-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M10 16v-1H3.5c-.83 0-1.5.67-1.5 1.5S2.67 18 3.5 18H10v-2zm10-8h-6V4c0-1.11-.89-2-2-2h-4c-1.11 0-2 .89-2 2v4H2v11c0 1.11.89 2 2 2h16c1.11 0 2-.89 2-2V8zm-8-4h4v4h-4V4zm6 13H6v-7h12v7z"/>
                </svg>`;
        } else if (name.includes('shop')) {
            return `
                <svg class="w-full h-full text-emerald-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M7 18c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zM1 2v2h2l3.6 7.59-1.35 2.45c-.16.28-.25.61-.25.96 0 1.1.9 2 2 2h12v-2H7.42c-.14 0-.25-.11-.25-.25l.03-.12L8.1 13h7.45c.75 0 1.41-.41 1.75-1.03L21.7 4H5.21l-.94-2H1zm16 16c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/>
                </svg>`;
        } else if (name.includes('lux')) {
            return `
                <svg class="w-full h-full text-yellow-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
                </svg>`;
        } else if (name.includes('forex')) {
            return `
                <svg class="w-full h-full text-teal-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M11.8 10.9c-2.27-.59-3-1.2-3-2.15 0-1.09 1.01-1.85 2.7-1.85 1.78 0 2.44.85 2.5 2.1h2.21c-.07-1.72-1.12-3.3-3.21-3.81V3h-3v2.16c-1.94.42-3.5 1.68-3.5 3.61 0 2.31 1.91 3.46 4.7 4.13 2.5.6 3 1.48 3 2.41 0 .69-.49 1.79-2.7 1.79-2.06 0-2.87-.92-2.98-2.1h-2.2c.12 2.19 1.76 3.42 3.68 3.83V21h3v-2.15c1.95-.37 3.5-1.5 3.5-3.55 0-2.84-2.43-3.81-4.7-4.4z"/>
                </svg>`;
        } else {
            // Default icon for unknown services
            return `
                <svg class="w-full h-full text-gray-600" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                </svg>`;
        }
    }

    getServiceColor(serviceName) {
        const name = serviceName.toLowerCase();
        
        if (name.includes('hotel') || name.includes('guesthouse')) return 'from-blue-500 to-blue-700';
        if (name.includes('visa')) return 'from-purple-500 to-purple-700';
        if (name.includes('flight')) return 'from-sky-500 to-sky-700';
        if (name.includes('bus')) return 'from-green-500 to-green-700';
        if (name.includes('car') || name.includes('ride')) return 'from-red-500 to-red-700';
        if (name.includes('love')) return 'from-pink-500 to-pink-700';
        if (name.includes('nomad') || name.includes('little')) return 'from-orange-500 to-orange-700';
        if (name.includes('event') || name.includes('retreat')) return 'from-indigo-500 to-indigo-700';
        if (name.includes('job')) return 'from-slate-500 to-slate-700';
        if (name.includes('shop')) return 'from-emerald-500 to-emerald-700';
        if (name.includes('lux')) return 'from-yellow-500 to-yellow-700';
        if (name.includes('forex')) return 'from-teal-500 to-teal-700';
        
        return 'from-gray-500 to-gray-700';
    }

    connectedCallback() {
        const serviceIcon = this.getServiceIcon(this.serviceType.name);
        const serviceColor = this.getServiceColor(this.serviceType.name);
        
        // Determine the link URL based on service type
        const linkUrl = this.getLinkUrl(this.serviceType.name);
        
        this.innerHTML = `
            <div class="group cursor-pointer h-full" onclick="window.location.href='${linkUrl}'">
                <div class="bg-white rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2 overflow-hidden border border-gray-100 h-full flex flex-col">
                    <div class="relative overflow-hidden bg-gradient-to-br from-gray-50 to-gray-100 flex items-center justify-center h-48">
                        <div class="w-20 h-20 bg-gradient-to-br ${serviceColor} rounded-2xl flex items-center justify-center shadow-lg group-hover:scale-110 transition-transform duration-500 p-4">
                            ${serviceIcon}
                        </div>
                        <div class="absolute inset-0 bg-gradient-to-t from-black/5 to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300"></div>
                        <div class="absolute top-4 right-4 w-10 h-10 bg-white/90 backdrop-blur-sm rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-all duration-300 transform translate-x-2 group-hover:translate-x-0">
                            <svg class="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
                            </svg>
                        </div>
                    </div>
                    <div class="p-6 flex-1 flex flex-col">
                        <h3 class="text-xl font-bold text-gray-900 mb-2 group-hover:text-blue-600 transition-colors duration-200">${this.serviceType.name}</h3>
                        <p class="text-gray-600 text-sm mb-4 flex-1">${this.serviceType.description || 'Discover amazing experiences and create unforgettable memories'}</p>
                        <div class="flex items-center justify-between mt-auto">
                            <span class="text-sm font-medium text-blue-600 bg-blue-50 px-3 py-1 rounded-full">Explore</span>
                            <div class="w-8 h-8 bg-gradient-to-r ${serviceColor} rounded-full flex items-center justify-center transform group-hover:scale-110 transition-transform duration-200">
                                <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6"></path>
                                </svg>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        `;
    }

    getLinkUrl(serviceName) {
        const name = serviceName.toLowerCase();
        
        if (name.includes('hotel') || name.includes('guesthouse')) {
            return '/hotels';
        } else if (name.includes('visa')) {
            return '/visa-assistance';
        } else if (name.includes('flight')) {
            return '/flights';
        } else if (name.includes('bus')) {
            return '/bus-travel';
        } else if (name.includes('car') || name.includes('ride')) {
            return '/car-rentals';
        } else if (name.includes('love')) {
            return '/nomado-love';
        } else if (name.includes('nomad') || name.includes('little')) {
            return '/little-nomads';
        } else if (name.includes('event') || name.includes('retreat')) {
            return '/events-retreats';
        } else if (name.includes('job')) {
            return '/nomado-jobs';
        } else if (name.includes('shop')) {
            return '/nomado-shop';
        } else if (name.includes('lux')) {
            return '/nomado-lux';
        } else if (name.includes('forex')) {
            return '/nomado-forex';
        } else {
            return '#'; // Default fallback
        }
    }
}

customElements.define("service-type-item", ServiceTypeItemComponent);