const countriesUrl = '/static/countries.json';
const countriesSelectId = [
  'nationality_1',
  'nationality_2',
]
const countriesCacheKey = 'countriesData';

const languagesUrl = '/static/languages.json';
const languagesSelectId = [
  'spoken_language_1',
  'spoken_language_2',  
  'spoken_language_3',  
  'preferred_communication_language',
] 
const languagesCacheKey = 'languagesData';

const cacheExpiry = 365 * 24 * 60 * 60 * 1000; // 1 year in milliseconds 

function populateSelect(selectId, data) {
    const selectElement = document.getElementById(selectId);
    data.forEach(item => {
        const option = document.createElement('option');
        option.value = item.value; // Adjust based on your JSON structure
        option.textContent = item.label; // Adjust based on your JSON structure
        selectElement.appendChild(option);
    });
}

function fetchData(dataUrl, cacheKey) {
    fetch(dataUrl)
        .then(response => response.json())
        .then(data => {
            localStorage.setItem(cacheKey, JSON.stringify({data: data, timestamp: Date.now()}));
            populateSelect(data);
        })
        .catch(error => console.error('Fetching data failed:', error));
}

function init(dataUrl, selectId, cacheKey) {
    const cached = localStorage.getItem(cacheKey);
    if (cached) {
        const {data, timestamp} = JSON.parse(cached);
        if (Date.now() - timestamp < cacheExpiry) {
            for (let i = 0; i < selectId.length; i++) {
                populateSelect(selectId[i], data);
            }
            return;
        }
    }
    fetchData(dataUrl, cacheKey);
}

function init_countries() {
    init(countriesUrl, countriesSelectId, countriesCacheKey);
}

function init_languages() {
    init(languagesUrl, languagesSelectId, languagesCacheKey);
}