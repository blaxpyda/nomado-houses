export const API = {
    baseURL: '/api/',
    getAllServices: async () => {
        return await API.fetch('api/service-types/');
    },
    // getServiceById: async (id) => {
    //     return await API.fetch(`api/services/${id}`);
    // },
    getAllDestinations: async () => {
        return await API.fetch('api/destinations/');
    },
    // getDestinationById: async (id) => {
    //     return await API.fetch(`api/destinations/${id}`);
    // },
    // searchDestinations: async (q, order, genre) => {
    //     return await API.fetch(`api/destinations/search/`, { q, order, genre });
    // },
    fetch: async (serviceName, args) => {
        try {
            const queryString = args ? new URLSearchParams(args).toString() : '';
            const response = await fetch(API.baseURL + serviceName);
            const result = await response.json();
            return result;
        } catch (error) {
            console.error(error);
        }
    }
}