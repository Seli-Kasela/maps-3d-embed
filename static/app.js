// Global variables
let map;
let mapConfig;
let userLocation;

/**
 * Initializes the application on page load
 */
async function initializeApp() {
    try {
        // Fetch geolocation data
        userLocation = await fetchGeolocation();
        displayLocationInfo(userLocation);

        // Fetch map configuration
        mapConfig = await fetchMapConfig();

        // Initialize the 3D map
        initializeMap();

        // Setup event listeners
        setupEventListeners();
    } catch (error) {
        console.error('Initialization error:', error);
        displayError('Failed to initialize the application');
    }
}

/**
 * Fetches user geolocation from the backend
 * @returns {Promise<Object>} Geolocation data
 */
async function fetchGeolocation() {
    const response = await fetch('/api/geolocation');
    if (!response.ok) {
        throw new Error('Failed to fetch geolocation');
    }
    return response.json();
}

/**
 * Fetches map configuration from the backend
 * @returns {Promise<Object>} Map configuration data
 */
async function fetchMapConfig() {
    const response = await fetch('/api/map-config');
    if (!response.ok) {
        throw new Error('Failed to fetch map configuration');
    }
    return response.json();
}

/**
 * Initializes the 3D Google Map
 */
function initializeMap() {
    const mapOptions = {
        center: {
            lat: mapConfig.location.latitude,
            lng: mapConfig.location.longitude,
        },
        zoom: mapConfig.zoom,
        tilt: mapConfig.tilt,
        heading: mapConfig.heading,
        mapTypeId: 'satellite', // Use satellite view for 3D effect
    };

    map = new google.maps.Map(document.getElementById('map'), mapOptions);

    // Add a marker at the location
    new google.maps.Marker({
        position: mapOptions.center,
        map: map,
        title: `${mapConfig.location.city}, ${mapConfig.location.country}`,
    });
}

/**
 * Displays location information on the info panel
 * @param {Object} location - Geolocation data
 */
function displayLocationInfo(location) {
    const infoPanel = document.getElementById('location-info');
    infoPanel.innerHTML = `
        <div>
            <p><strong>IP Address:</strong><br>${location.ip}</p>
            <p><strong>City:</strong><br>${location.city}</p>
            <p><strong>Country:</strong><br>${location.country}</p>
            <p><strong>Coordinates:</strong><br>${location.latitude.toFixed(4)}, ${location.longitude.toFixed(4)}</p>
            <p><strong>ISP:</strong><br>${location.isp}</p>
        </div>
    `;
}

/**
 * Displays error message to the user
 * @param {string} message - Error message
 */
function displayError(message) {
    const infoPanel = document.getElementById('location-info');
    infoPanel.innerHTML = `<p style="color: #e74c3c;"><strong>Error:</strong> ${message}</p>`;
}

/**
 * Sets up event listeners for user interactions
 */
function setupEventListeners() {
    const refreshBtn = document.getElementById('refresh-btn');
    refreshBtn.addEventListener('click', handleRefresh);
}

/**
 * Handles the refresh button click
 */
async function handleRefresh() {
    try {
        const refreshBtn = document.getElementById('refresh-btn');
        refreshBtn.disabled = true;
        refreshBtn.textContent = 'Loading...';

        // Refresh geolocation data
        userLocation = await fetchGeolocation();
        displayLocationInfo(userLocation);

        // Update map position if location changed
        if (map) {
            const newCenter = {
                lat: userLocation.latitude,
                lng: userLocation.longitude,
            };
            map.setCenter(newCenter);
        }

        refreshBtn.textContent = 'Refresh Location';
        refreshBtn.disabled = false;
    } catch (error) {
        console.error('Refresh error:', error);
        displayError('Failed to refresh location data');

        const refreshBtn = document.getElementById('refresh-btn');
        refreshBtn.textContent = 'Refresh Location';
        refreshBtn.disabled = false;
    }
}

/**
 * Rotates the map view
 * @param {number} angle - Rotation angle in degrees
 */
function rotateMap(angle) {
    if (map) {
        map.setHeading(angle);
    }
}

/**
 * Tilts the map view
 * @param {number} tilt - Tilt angle in degrees
 */
function tiltMap(tilt) {
    if (map) {
        map.setTilt(tilt);
    }
}

/**
 * Zooms the map
 * @param {number} level - Zoom level
 */
function zoomMap(level) {
    if (map) {
        map.setZoom(level);
    }
}

// Initialize the application when the DOM is ready
if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', initializeApp);
} else {
    initializeApp();
}
