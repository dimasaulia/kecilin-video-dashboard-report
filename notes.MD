# Record Report

## Test Case

Terdapat rekaman video dari cctv dengan durasi potongan per-file videonya adalah 5 menit, terkadang terjadi error/stop record ataupun karena faktor eksternal (ex: mati listrik) maka akan ada waktu dimana tidak ada file rekamannya, dari persoalan tadi client perlu report harian yang isinya adalah data berapa lama total rekaman di hari itu, berapa lama total error record pada hari itu, jumlah keduanya (total record dan total error) harus sama dengan 1 hari (24 jam).
buatkan record report yang berisi:

```sh
{
    "start_time": Filter waktu mulai,
    "end_time": Filter waktu selesai,
    "result": {
        "error_percentage": Nilai presentase error,
        "error_time": Total durasi error (detik),
        "record_percentage": Nilai presentase recording,
        "record_time": Total durasi recording (detik),
        "total_error": Jumlah titik error,
        "total_recording": Jumlah file record,
        "total_time": Total waktu (start_time ~ end_time) (24 Jam),
        "error": [ # Detail data error
            {
                "duration": Durasi error (detik),
                "filename": Nama file terakhir sebelum error recording (kalau tidak ada Null),
                "time_error": { # Detail waktu error
                    "end_time": Waktu selesai error,
                    "start_time": Waktu mulai error
                }
            }
        ],
        "recording_file": [ # Detail file record
            {
                "duration": Durasi file,
                "filename": Nama file
            }
        ]
    }
}
```

Contoh :

```sh
{
    "start_time": "2023-12-14 00:00:00",
    "end_time": "2023-12-15 00:00:00",
    "result": {
        "error_percentage": 82.51736111111111,
        "error_time": 71295,
        "record_percentage": 17.48263888888889,
        "record_time": 15105,
        "total_error": 2,
        "total_recording": 53,
        "total_time": 86400,
        "error": [
            {
                "duration": 30,
                "filename": "2023-12-14T14-18-57.mp4",
                "time_error": {
                    "end_time": "2023-12-14 14:19:32",
                    "start_time": "2023-12-14 14:19:02"
                }
            },
            .
            .
            .
            {
                "duration": 600,
                "filename": "2023-12-14T16-59-32.mp4",
                "time_error": {
                    "end_time": "2023-12-14 17:10:12",
                    "start_time": "2023-12-14 17:00:12"
                }
            }
        ],
        "recording_file": [
            {
                "duration": 300,
                "filename": "2023-12-14T13-03-57.mp4"
            },
            {
                "duration": 300,
                "filename": "2023-12-14T13-08-57.mp4"
            },
            .
            .
            .
            .
            {
                "duration": 40,
                "filename": "2023-12-14T16-59-32.mp4"
            }
        ]
    }
}
```

\
\
Berdasarkan sample file rekaman yang diberikan, buatkan record report pada hari:

1. start_date = 2024-01-09 00-00-00, end_date=2024-01-10 00-00-00
2. start_date = 2024-01-10 00-00-00, end_date=2024-01-11 00-00-00
3. start_date = 2024-01-11 00-00-00, end_date=2024-01-12 00-00-00

Upload project ke gitlab / github / etc (code dan result record report) lalu share link-nya ke email: rosyidridho@kecilin.id, cc: gayarijamila@kecilin.id
