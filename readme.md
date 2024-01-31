# Kecilin Video Dahboard

---

Sistem di kembangkan menggunakan golang dengan library fiber untuk menangani http request. Untuk menjalankan progman ini cukup dengan perintah.

```sh
go run main.go
```

Sebelum menjalankan program tersebut atur configurasi server melalui .env. Buat file env dengan perintah `nano .env` lalu isi file tersebut dengan data berikut

```sh
PORT = 8000 (isi dengan angka)
FORK = false (isi dengan boolean)
```

Logic atau fungsi utama yang melakukan proses scanning dapat ditemukan pada file [berikut ini](/app/api/video_record/service/vid_rec_svc_impl.go)

### User interface

Setelah mengaktifkan server maka anda dapat mengunjungi halaman dashboard untuk melihat presentasi keberhasilan record beserta melihat file yang mengalami kerusakan.

![Video](/doc/dashboard.gif)

Berikut adalah hasil respon dari api yang berfungsi untuk menampilkan report dari video directory.

### Periode: 2024-01-09T00-00-00 hingga 2024-01-10T00-00-00

Url: `/api/v1/video/report?start_date=2024-01-09T00-00-00&end_date=2024-01-10T00-00-00`

Response:

```sh
{
    "start_time": "2024-01-09T00:00:00Z",
    "end_time": "2024-01-10T00:00:00Z",
    "result": {
        "error_percentage": 1.7361111111111112,
        "error_time": 1500,
        "record_percentage": 98.26388888888889,
        "record_time": 84900,
        "total_error": 2,
        "total_recording": 283,
        "total_time": 86400,
        "error": [
            {
                "duration": 300,
                "filename": "2024-01-09T05-38-00.mp4",
                "TimeError": {
                    "end_time": "2024-01-09T05:43:00Z",
                    "start_time": "2024-01-09T05:38:00Z"
                }
            },
            {
                "duration": 1200,
                "filename": "2024-01-09T16-08-00.mp4",
                "TimeError": {
                    "end_time": "2024-01-09T16:28:00Z",
                    "start_time": "2024-01-09T16:08:00Z"
                }
            }
        ],
        "recording_file": [
            {
                "duration": 300,
                "filename": "2024-01-09T00-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T00-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T01-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T02-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T03-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T04-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T05-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T06-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T07-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T08-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T09-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T10-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T11-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T12-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T13-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T14-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T15-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T16-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T17-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T18-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T19-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T20-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T21-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T22-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-09T23-53-00.mp4"
            }
        ]
    }
}
```

### Periode: 2024-01-10T00-00-00 hingga 2024-01-11T00-00-00

Url: `/api/v1/video/report?start_date=2024-01-10T00-00-00&end_date=2024-01-11T00-00-00`

Response:

```sh
{
    "start_time": "2024-01-10T00:00:00Z",
    "end_time": "2024-01-11T00:00:00Z",
    "result": {
        "error_percentage": 3.125,
        "error_time": 2700,
        "record_percentage": 96.875,
        "record_time": 83700,
        "total_error": 4,
        "total_recording": 279,
        "total_time": 86400,
        "error": [
            {
                "duration": 900,
                "filename": "2024-01-10T06-03-00.mp4",
                "TimeError": {
                    "end_time": "2024-01-10T06:18:00Z",
                    "start_time": "2024-01-10T06:03:00Z"
                }
            },
            {
                "duration": 1200,
                "filename": "2024-01-10T12-08-00.mp4",
                "TimeError": {
                    "end_time": "2024-01-10T12:28:00Z",
                    "start_time": "2024-01-10T12:08:00Z"
                }
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-58-00.mp4",
                "TimeError": {
                    "end_time": "2024-01-10T20:03:00Z",
                    "start_time": "2024-01-10T19:58:00Z"
                }
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-33-00.mp4",
                "TimeError": {
                    "end_time": "2024-01-10T22:38:00Z",
                    "start_time": "2024-01-10T22:33:00Z"
                }
            }
        ],
        "recording_file": [
            {
                "duration": 300,
                "filename": "2024-01-10T00-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T00-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T01-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T02-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T03-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T04-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T05-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T06-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T07-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T08-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T09-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T10-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T11-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T12-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T13-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T14-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T15-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T16-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T17-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T18-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T19-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T20-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T21-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T22-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-10T23-53-00.mp4"
            }
        ]
    }
}
```

### Periode: 2024-01-11T00-00-00 hingga 2024-01-12T00-00-00

Url: `/api/v1/video/report?start_date=2024-01-11T00-00-00&end_date=2024-01-12T00-00-00`

Response:

```sh
{
    "start_time": "2024-01-11T00:00:00Z",
    "end_time": "2024-01-12T00:00:00Z",
    "result": {
        "error_percentage": 0,
        "error_time": 0,
        "record_percentage": 100,
        "record_time": 86400,
        "total_error": 0,
        "total_recording": 288,
        "total_time": 86400,
        "error": [],
        "recording_file": [
            {
                "duration": 300,
                "filename": "2024-01-11T00-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T00-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T01-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T02-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T03-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T04-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T05-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T06-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T07-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T08-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T09-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T10-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T11-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T12-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T13-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T14-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T15-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T16-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T17-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T18-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T19-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T20-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T21-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T22-58-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-03-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-08-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-13-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-18-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-23-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-28-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-33-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-38-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-43-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-48-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-53-00.mp4"
            },
            {
                "duration": 300,
                "filename": "2024-01-11T23-53-00.mp4"
            }
        ]
    }
}
```
