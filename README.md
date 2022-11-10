# Valūtas kurss iekš MySQL, izmantojot Go programmēšanas valodu - Matīss Petrausks

## Sagataves programmatūras
```
- Docker
- MySQL
- Go programmēšanas valoda
```

## Kā kļūt Valūtas kursu datubāzē, izmantojot Docker platformu?

Sākumā, lietotājam ir jāsakonfigurē Docker platforma un MySQL serveris, un jāuzsāk savienošanās, izmantojot lejā dotos pievienošanās datus -
```
Parole - parole
Datubāžu vārds - valutaKurss

docker run ^
 --rm ^
 -d ^
 -e MYSQL_ROOT_PASSWORD=parole ^
 -e MYSQL_DATABASE=valutaKurss ^
 -p 3306:3306 ^
 mysql:8.0.27
```
Lietotājam pēc tā tiks dots konteinera kods, kuru būs jāievada, lai pievienotos specifiski valutaKurss datubāzē.
```
docker exec -it {jūsu konteinera kods} mysql -u root -p -h localhost
Enter password: parole
```
Pēc tā, lietotājs atrastos precīzi savstarpējā datubāzē valutaKurss, pēc tā datubāze ir jāuzbūvē -
```
CREATE TABLE channel (
  id int(200) NOT NULL AUTO_INCREMENT,
  title varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  description text COLLATE utf8_unicode_ci NOT NULL,
  pubDate varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO channel (id, title, description, pubDate) VALUES
(1, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'AUD 1.54280000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(2, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'BGN 1.95580000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(3, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'BRL 5.07000000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(4, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'CAD 1.34640000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(5, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'CHF 0.98740000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(6, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'CNY 7.21890000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(7, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'CZK 24.30100000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(8, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'DKK 7.43930000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(9, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'GBP 0.87135000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(10, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'HKD 7.84440000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(11, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'HRK 7.53750000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(12, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'HUF 401.03000000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(13, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'IDR 15648.95000000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(14, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'ILS 3.54020000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(15, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'INR 81.84070000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(16, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'ISK 145.90000000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(17, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'JPY 146.18000000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(18, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'KRW 1391.25000000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(19, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'MXN 19.43950000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(20, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'MYR 4.73620000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(21, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'NOK 10.25550000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(22, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'NZD 1.68340000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(23, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'PHP 58.36100000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(24, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'PLN 4.68650000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(25, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'RON 4.88550000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(26, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'SEK 10.83200000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(27, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'SGD 1.40220000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(28, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'THB 37.28400000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(29, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'TRY 18.58750000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(30, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'USD 0.99930000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(31, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'ZAR 17.75830000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(32, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 7. November', 'MYR 4.73620000', 'Mon, 07 Nov 2022 02:00:00 +0200'),
(33, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'AUD 1.54350000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(34, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'BGN 1.95580000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(35, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'BRL 5.20300000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(36, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'CAD 1.34890000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(37, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'CHF 0.99110000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(38, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'CNY 7.24950000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(39, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'CZK 24.32600000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(40, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'DKK 7.43780000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(41, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'GBP 0.87378000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(42, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'HKD 7.84680000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(43, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'HRK 7.53900000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(44, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'HUF 400.75000000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(45, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'IDR 15652.76000000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(46, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'ILS 3.54360000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(47, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'INR 81.51800000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(48, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'ISK 146.30000000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(49, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'JPY 146.25000000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(50, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'KRW 1377.94000000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(51, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'MXN 19.44950000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(52, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'MYR 4.73460000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(53, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'NOK 10.27950000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(54, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'NZD 1.68600000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(55, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'PHP 58.18700000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(56, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'PLN 4.69180000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(57, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'RON 4.89780000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(58, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'SEK 10.83730000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(59, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'SGD 1.40220000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(60, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'THB 37.22000000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(61, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'TRY 18.59910000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(62, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'USD 0.99930000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(63, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 8. November', 'ZAR 17.83970000', 'Mon, 08 Nov 2022 02:00:00 +0200'),
(64, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'AUD 1.55380000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(65, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'BGN 1.95580000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(66, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'BRL 5.19470000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(67, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'CAD 1.35010000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(68, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'CHF 0.98800000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(69, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'CNY 7.28130000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(70, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'CZK 24.33700000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(71, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'DKK 7.43820000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(72, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'GBP 0.87774000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(73, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'HKD 7.88010000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(74, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'HRK 7.54250000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(75, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'HUF 403.53000000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(76, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'IDR 15717.07000000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(77, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'ILS 3.56210000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(78, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'INR 81.65750000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(79, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'ISK 146.70000000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(80, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'JPY 146.82000000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(81, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'KRW 1369.73000000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(82, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'MXN 19.65540000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(83, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'MYR 4.70980000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(84, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'NOK 10.32200000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(85, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'NZD 1.70330000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(86, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'PHP 58.23600000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(87, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'PLN 4.70100000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(88, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'RON 4.90450000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(89, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'SEK 10.84500000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(90, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'SGD 1.40610000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(91, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'THB 36.99900000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(92, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'TRY 18.67280000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(93, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'USD 1.00390000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(94, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 9. November', 'ZAR 17.87700000', 'Mon, 09 Nov 2022 02:00:00 +0200'),
(95, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'AUD 1.55250000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(96, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'BGN 1.95580000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(97, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'BRL 5.28600000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(98, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'CAD 1.34670000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(99, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'CHF 0.98340000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(100, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'CNY 7.21840000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(101, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'CZK 24.36100000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(102, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'DKK 7.43810000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(103, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'GBP 0.87298000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(104, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'HKD 7.81280000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(105, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'HRK 7.54270000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(106, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'HUF 400.95000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(107, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'IDR 15615.60000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(108, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'ILS 3.54530000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(109, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'INR 81.30580000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(110, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'ISK 147.50000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(111, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'JPY 145.47000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(112, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'KRW 1373.96000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(113, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'MXN 19.45620000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(114, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'MYR 4.67890000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(115, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'NOK 10.36150000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(116, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'NZD 1.69840000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(117, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'PHP 57.79300000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(118, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'PLN 4.70600000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(119, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'RON 4.89130000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(120, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'SEK 10.87430000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(121, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'SGD 1.39630000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(122, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'THB 36.70000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(123, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'TRY 18.51000000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(124, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'USD 0.99540000', 'Mon, 10 Nov 2022 02:00:00 +0200'),
(125, 'Eiropas Centrālās bankas publicētie eiro atsauces kursi. 10. November', 'ZAR 17.68820000', 'Mon, 10 Nov 2022 02:00:00 +0200');
```
Uzbūvējot datubāzi, ir pieejami dati, kurus var apstrādāt kā lietotājs vēlas, kā piemēram datubāzē visjaunākos valūtas kursus un specifisku iepriekšējo kursu kādai valstij - 
```
SELECT title, description, pubDate
from channel
where pubdate like "%10 Nov 2022%";

SELECT title, description, pubDate
from channel
where pubdate like "%07 Nov 2022%" and description like "USD%";
```
Ja lietotājs vēlas saglabāt datubāzes kopiju, var izmantot šo komandu, ir jāiziet no MySQL un jānorāda datoram atrašanās vietu, kur lietotājs vēlas saglabāt datubāzi - 
```
cd C:\Users\{Lietotājs}\Desktop\Valutu kursi un vertibas

mysqldump -u root -p valutaKurss > valutaKurss.sql
```