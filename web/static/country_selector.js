class CountrySelector {
    constructor(
        inputId,
        countryListId,
        autocompleteId,
        countries = [],
        selectedCountryCodes = [],
    ) {
        this.inputID = inputId
        this.coutryListId = countryListId
        this.autocompleteId = autocompleteId
        this.countriesSubject = new rxjs.BehaviorSubject(countries)
        this.countries$ = this.countriesSubject.asObservable()
        this.selectedCountryCodesSubject = new rxjs.BehaviorSubject(selectedCountryCodes)
        this.selectedCountryCodes$ = this.selectedCountryCodesSubject.asObservable()
        this.selectedCountries$ = this.selectedCountryCodes$.pipe(
            rxjs.withLatestFrom(this.countries$),
            rxjs.map(([codes, countries]) => {
                return countries.filter(c => codes.includes(c.value))
            })
        )
        this.unselectedCountries$ = this.selectedCountryCodes$.pipe(
            rxjs.withLatestFrom(this.countries$),
            rxjs.map(([codes, countries]) => {
                return countries.filter(c => !codes.includes(c.value))
            })
        )

        this.availableCountries$ = this.unselectedCountries$.pipe(
            rxjs.map(countries => {
                return countries.filter(c => !c.readOnly)
            })
        )


        this.setupEvents()
    }

    selectCountry(code) {
        const country = this.countriesSubject.value.find(c => c.value === code)
        if (!country) {
            return
        }
        if (this.selectedCountryCodesSubject.value.includes(code)) {
            return
        }
        const newValue = [...this.selectedCountryCodesSubject.value, code]
        newValue.sort((a, b) => a.localeCompare(b))
        this.selectedCountryCodesSubject.next(newValue)
    }

    removeCountry(code) {
        if (this.selectedCountryCodesSubject.value.includes(code)) {
            this.selectedCountryCodesSubject.next(this.selectedCountryCodesSubject.value.filter(x => x !== code))
        }
    }

    setupEvents() {
        const init = this.init.bind(this)
        document.addEventListener("DOMContentLoaded", init)
    }

    init() {
        const $this = this
        $this.selectedCountries$.subscribe(countries => {

            const countryInput = document.getElementById($this.inputID)
            countryInput.value = countries.map(c => c.value).join(",")

            const countryList = document.getElementById($this.coutryListId)
            countryList.innerHTML = ""

            if (countries.length === 0) {
                console.log("no countries selected")
                const emptyItem = document.createElement("div")
                emptyItem.className = "list-group-item"
                emptyItem.innerHTML = "No countries selected"
                countryList.appendChild(emptyItem)
                return
            }

            countries.forEach(function ({value, label, readOnly}) {
                const countryItem = document.createElement("div")
                countryItem.className = "list-group-item"

                const labelSpan = document.createElement("span")
                labelSpan.innerHTML = label
                countryItem.appendChild(labelSpan)

                console.log("readonly", readOnly)
                if (!readOnly) {
                    const buttonSpan = document.createElement("span")
                    buttonSpan.className = "badge text-dark bg-light float-end"
                    buttonSpan.style.cursor = "pointer"

                    const removeCountry = $this.removeCountry.bind($this, value)
                    buttonSpan.onclick = () => removeCountry()
                    buttonSpan.innerHTML = "&times;"

                    countryItem.appendChild(buttonSpan)
                } else {
                    countryItem.className += " text-secondary"
                }

                countryList.appendChild(countryItem)
            })

        })

        const countryInput = document.getElementById(this.autocompleteId)
        const ac = new Autocomplete(countryInput, {
            data: [],
            threshold: 0,
            maximumItems: 5,
            onSelectItem: ({label, value}) => {
                this.selectCountry(value)
                if (countryInput.value !== "") {
                    countryInput.value = ""
                }
            }
        })

        this.availableCountries$.pipe(
            rxjs.distinctUntilChanged((a, b) => JSON.stringify(a) === JSON.stringify(b)),
        ).subscribe(countries => {
            ac.setData(countries)
            ac.dropdown.hide()
            countryInput.disabled = countries.length === 0;
        })
    }
}
