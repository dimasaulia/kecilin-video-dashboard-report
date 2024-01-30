const datepickerStart = document.getElementById("start-date");
const datepickerEnd = document.getElementById("end-date");
const btn = document.getElementById("search-button");
const durationBadPlaceholder = document.getElementById("duration_bad");
const durationGoodPlaceholder = document.getElementById("duration_good");
const durationTotalPlaceholder = document.getElementById("duration_total");
const fileBadPlaceholder = document.getElementById("file_bad");
const fileGoodPlaceholder = document.getElementById("file_good");
const fileTotalPlaceholder = document.getElementById("file_total");
const vidContainer = document.getElementById("vid-container");

function videoTemplate(file) {
    return `
    <div class="col-span-1" >
        <div class="max-w-sm rounded-custom bg-white border border-gray-200 rounded-lg shadow-smooth overflow-hidden dark:bg-gray-800 dark:border-gray-700">
                <video class="h-40 lazy" controls playsinline muted preload="metadata" poster="/public/image/placeholder.png">
                    <source src="/public/video/${file}" type="video/mp4">
                    Your browser does not support the video tag.
                </video>
            <div class="p-3">
                <p>${file}</p>
            </div>
        </div>              
    </div>
    `;
}

function transformDateFormat(inputDate) {
    // Parse the input date string
    let dateParts = inputDate.split("/");
    let parsedDate = new Date(
        `${dateParts[2]}-${dateParts[0]}-${dateParts[1]}`
    );

    // Format the date into the desired format
    let formattedDate = `${parsedDate.getFullYear()}-${padZero(
        parsedDate.getMonth() + 1
    )}-${padZero(parsedDate.getDate())}T00-00-00`;

    return formattedDate;
}

function padZero(value) {
    return value < 10 ? "0" + value : value;
}

String.prototype.replaceAt = function (index, char) {
    let a = this.split("");
    a[index] = char;
    return a.join("");
};

function bubbleSort(arr) {
    const n = arr.length;

    for (let i = 0; i < n - 1; i++) {
        for (let j = 0; j < n - 1 - i; j++) {
            // Compare adjacent elements
            if (compareDates(arr[j].filename, arr[j + 1].filename) > 0) {
                // Swap elements if they are in the wrong order
                const temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
            }
        }
    }

    return arr;
}

function compareDates(dateA, dateB) {
    const parsedDateA = new Date(
        dateA.replaceAt(13, ":").replaceAt(16, ":").replace(".mp4", "")
    );
    const parsedDateB = new Date(
        dateB.replaceAt(13, ":").replaceAt(16, ":").replace(".mp4", "")
    );
    return parsedDateA - parsedDateB;
}

// Render Dounat
let chart;
let newArr;
const getChartOptions = (badFile, goodFile) => {
    return {
        series: [badFile, goodFile],
        colors: ["#1C64F2", "#16BDCA"],
        chart: {
            height: 420,
            width: "100%",
            type: "pie",
        },
        stroke: {
            colors: ["white"],
            lineCap: "",
        },
        plotOptions: {
            pie: {
                labels: {
                    show: true,
                },
                size: "100%",
                dataLabels: {
                    offset: -25,
                },
            },
        },
        labels: ["Bad File", "Good File"],
        dataLabels: {
            enabled: true,
            style: {
                fontFamily: "Inter, sans-serif",
            },
        },
        legend: {
            position: "bottom",
            fontFamily: "Inter, sans-serif",
        },
        yaxis: {
            labels: {
                formatter: function (value) {
                    return value + "%";
                },
            },
        },
        xaxis: {
            labels: {
                formatter: function (value) {
                    return value + "%";
                },
            },
            axisTicks: {
                show: false,
            },
            axisBorder: {
                show: false,
            },
        },
    };
};

window.addEventListener("load", function () {
    if (
        document.getElementById("pie-chart") &&
        typeof ApexCharts !== "undefined"
    ) {
        chart = new ApexCharts(
            document.getElementById("pie-chart"),
            getChartOptions(50, 50)
        );
        chart.render();
    }
});

datepickerStart.value = "01/09/2024";
datepickerEnd.value = "01/10/2024";

function reportLoader(data) {
    let times = [];

    durationBadPlaceholder.textContent = data.result.error_time;
    durationGoodPlaceholder.textContent = data.result.record_time;
    durationTotalPlaceholder.textContent = data.result.total_time;
    fileBadPlaceholder.textContent = data.result.total_error;
    fileGoodPlaceholder.textContent = data.result.total_recording;
    fileTotalPlaceholder.textContent =
        Number(data.result.total_error) + Number(data.result.total_recording);

    const allFile = [];
    allFile.push(...data.result.recording_file);
    allFile.push(...data.result.error);
    newArr = bubbleSort(allFile);

    chart.updateOptions(
        getChartOptions(
            data.result.error_percentage,
            data.result.record_percentage
        )
    );

    for (let i = 0; i < newArr.length; i++) {
        const d = newArr[i];
        const nextD = newArr[i + 1];

        const currentTime = new Date(
            d.filename.replaceAt(13, ":").replaceAt(16, ":").replace(".mp4", "")
        ).getTime();

        const nextTime = new Date(
            nextD?.filename
                .replaceAt(13, ":")
                .replaceAt(16, ":")
                .replace(".mp4", "")
        ).getTime();

        if (d?.TimeError) {
            times.push({
                color: "red",
                starting_time: currentTime,
                ending_time: nextTime,
            });
        } else {
            times.push({
                color: "green",
                starting_time: currentTime,
                ending_time: nextTime || Number(currentTime) + 300_000,
            });
        }
    }

    var dataTimeLine = [
        {
            label: "",
            times: times,
        },
    ];

    var timelineChart = d3.timelines();
    document.getElementById("timeline1").textContent = "";
    var svg = d3
        .select("#timeline1")
        .append("svg")
        .attr("width", 1000)
        .datum(dataTimeLine)
        .call(timelineChart);

    vidContainer.textContent = "";
    data.result.error.forEach((d) => {
        vidContainer.insertAdjacentHTML("beforeend", videoTemplate(d.filename));
    });
}

btn.addEventListener("click", (e) => {
    e.preventDefault();
    const startDate = transformDateFormat(datepickerStart.value.trim());
    const endDate = transformDateFormat(datepickerEnd.value.trim());

    dataLoader({
        url: `/api/v1/video/report?start_date=${startDate}&end_date=${endDate}`,
        func: reportLoader,
    });
});

dataLoader({
    url: `/api/v1/video/report?start_date=2024-01-09T00-00-00&end_date=2024-01-10T00-00-00`,
    func: reportLoader,
});
