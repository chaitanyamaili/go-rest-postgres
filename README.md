# rest postgres

## curl command

POSY commands to insert into `/albums`
```
curl -X POST http://localhost:8080/albums -H "Content-Type: application/json"  -d '{
"id":"1",
"title":"title",
"artist":"artist name",
"price":100.00
}'
```

POSY commands to insert into `/mf_data`
```
curl -X POST http://localhost:8080/mf_data -H "Content-Type: application/json"  -d '{
"id":"2",
"fund_house":"fund_house2",
"scheme_type":"scheme_type2",
"scheme_category":"scheme_category2",
"scheme_code":"scheme_code2",
"scheme_name":"scheme_name2",
"date":"28-03-2024",
"nav":"12"
}'
```

pgAdmin host: `host.docker.internal`

| Mutual Fund | MF API URL | MF API LATEST URL |
|-------------|------------|------------  |
|Mirae Asset Great Consumer Fund Direct Growth| https://api.mfapi.in/mf/118837 | https://api.mfapi.in/mf/118837/latest |
|SBI Magnum Children's Benefit Fund - Investment Plan - Direct Plan| https://api.mfapi.in/mf/148490 | https://api.mfapi.in/mf/148490/latest |
|Axis Small Cap Fund - Direct Plan - Growth| https://api.mfapi.in/mf/125354 | https://api.mfapi.in/mf/125354/latest |
|Axis Bluechip Fund - Direct Plan - Growth| https://api.mfapi.in/mf/120465 | https://api.mfapi.in/mf/120465/latestx |
|HDFC Mid-Cap Opportunities Fund - Direct Plan - Growth| https://api.mfapi.in/mf/118989 | https://api.mfapi.in/mf/118989/latest
|Canara Robeco Bluechip Equity Fund - Direct Plan - Growth| https://api.mfapi.in/mf/118269 | https://api.mfapi.in/mf/118269/latest |
