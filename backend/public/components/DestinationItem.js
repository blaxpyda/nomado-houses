export class DestinationItemComponent extends HTMLElement {
    constructor(destination) {
        super();
        this.destination = destination;
    }
    connectedCallback() {
        this.innerHTML = `
            <div class="group cursor-pointer h-full">
                <div class="bg-white rounded-2xl shadow-lg hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-2 overflow-hidden border border-gray-100 h-full flex flex-col">
                    <!-- Image Section -->
                    <div class="relative overflow-hidden h-48">
                        <img src="${this.destination.image}" alt="${this.destination.name}" 
                             class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500">
                        <div class="absolute inset-0 bg-gradient-to-t from-black/40 via-transparent to-transparent"></div>
                        
                        <!-- Status Badge -->
                        <div class="absolute top-3 left-3 bg-green-500 text-white px-2 py-1 rounded-full text-xs font-medium flex items-center space-x-1">
                            <svg class="w-3 h-3" fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
                            </svg>
                            <span>Available</span>
                        </div>
                        
                        <!-- Wishlist Heart -->
                        <div class="absolute top-3 right-3 w-9 h-9 bg-white/20 backdrop-blur-sm rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-all duration-300 hover:bg-white/30 cursor-pointer">
                            <svg class="w-4 h-4 text-white hover:text-red-400 transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"></path>
                            </svg>
                        </div>
                        
                        <!-- Price Badge -->
                        <div class="absolute bottom-3 right-3 bg-white/95 backdrop-blur-sm px-3 py-1 rounded-full">
                            <div class="text-xs text-gray-600 font-medium">From</div>
                            <div class="text-lg font-bold text-gray-900">${this.destination.price || '$299'}</div>
                        </div>
                    </div>
                    
                    <!-- Content Section -->
                    <div class="p-5 flex-1 flex flex-col">
                        <!-- Header -->
                        <div class="mb-3">
                            <h3 class="text-xl font-bold text-gray-900 mb-1 group-hover:text-blue-600 transition-colors">${this.destination.name}</h3>
                            <p class="text-gray-600 text-sm line-clamp-2 leading-relaxed">${this.destination.description || 'Experience the beauty and culture of this amazing destination with unforgettable memories.'}</p>
                        </div>
                        
                        <!-- Rating and Reviews -->
                        <div class="flex items-center mb-4">
                            <div class="flex items-center space-x-1 mr-3">
                                <svg class="w-4 h-4 text-yellow-400 fill-current" viewBox="0 0 20 20">
                                    <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"></path>
                                </svg>
                                <span class="text-sm font-semibold text-gray-900">${this.destination.rating || '4.8'}</span>
                            </div>
                            <span class="text-sm text-gray-500">(${this.destination.reviews || '2.1k'} reviews)</span>
                        </div>
                        
                        <!-- Features/Highlights -->
                        <div class="flex flex-wrap gap-2 mb-4">
                            <span class="inline-block bg-blue-50 text-blue-700 text-xs font-medium px-2 py-1 rounded-full">Popular</span>
                            <span class="inline-block bg-green-50 text-green-700 text-xs font-medium px-2 py-1 rounded-full">Best Value</span>
                            ${this.destination.featured ? '<span class="inline-block bg-purple-50 text-purple-700 text-xs font-medium px-2 py-1 rounded-full">Featured</span>' : ''}
                        </div>
                        
                        <!-- Action Button -->
                        <div class="mt-auto">
                            <button class="w-full py-3 px-4 bg-gradient-to-r from-blue-600 to-purple-600 text-white font-semibold rounded-xl hover:from-blue-700 hover:to-purple-700 transform hover:scale-105 transition-all duration-200 shadow-lg hover:shadow-xl">
                                <span class="flex items-center justify-center space-x-2">
                                    <span>View Details</span>
                                    <svg class="w-4 h-4 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
                                    </svg>
                                </span>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        `;
    }
}
customElements.define('destination-item', DestinationItemComponent);