import { reactive } from "vue";

export const store = reactive({
    selectedTime: "2021-02",
    // month: "02",
    month: 2,
    // monthNum: 2,
    monthTitle: "01 January",
    year: 2021,
    firstYear: 2021,
    firstmonth: 4,
    thisYear: 2023,
    thismonth: 12,
    monthsList: [
        { title: "01 January", value: 1, enabled: true },
        { title: "02 February", value: 2, enabled: true },
        { title: "03 March", value: 3, enabled: true },
        { title: "04 April", value: 4, enabled: true },
        { title: "05 May", value: 5, enabled: true },
        { title: "06 June", value: 6, enabled: true },
        { title: "07 July", value: 7, enabled: true },
        { title: "08 August", value: 8, enabled: true },
        { title: "09 September", value: 9, enabled: true },
        { title: "10 October", value: 10, enabled: true },
        { title: "11 November", value: 11, enabled: true },
        { title: "12 December", value: 12, enabled: true },
    ],
    updateSelectedTime(t) {
        this.selectedTime = t
    },
    getMonth() {
        // return
        let m = (this.month).toString()
        return m.length < 2 ? `0${m}` : `${m}`
    }
})
