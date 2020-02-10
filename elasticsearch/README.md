Testing Ground for ElasticSearch in search of suitable index model.

The data is a simple id and attributes.
The challenge is to find documents with specific attribute (key & value) and to find the aggregate of attributes value based on given attribute key.

Data generator use specific seed so it can generate the same data for each models.


## Running ElasticSearch

Run an elastic-search server using docker: 

```
docker run --name estest68 --rm -p 9200:9200 blacktop/elasticsearch:6.8
docker run --name estest75 --rm -p 9200:9200 blacktop/elasticsearch:7.5
```

Using ElasticSearch 7, if you want to count real total hits, add this in url parameter: `?track_total_hits=true`.
Ref: https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-track-total-hits

## Test Comparison
- Take 1, use nested document (key & value) for attributes.
  
  ```json
  {
    "attributes":[
      {"key":"Color", "value":"Orange"},
      {"key":"Size", "value":"XL"}
    ]
  }
  ```

- Take 2, use flat document by concatenating attributes key & value.

  ```json
  {
    "attributes":["Color:Orange", "Size:XL"]
  }
  ```

## Test Run
1. Find document with `Color`:`Orange` and `Size`:`XL`
2. Find number of document for each color
3. Find unique attributes.key (should return `Color` and `Size`)
4. Find number of document for each size for Orange color
5. Find Orange XL, show aggregate for Color (only for XL) & and Sizes (only for Orange) 

```bash
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take1/_search -d '@take1/orange_xl.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take1/_search -d '@take1/colors.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take1/_search -d '@take1/keys.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take1/_search -d '@take1/orange_sizes.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take1/_search -d '@take1/orange_xl_facet.json'

curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take2/_search -d '@take2/orange_xl.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take2/_search -d '@take2/colors.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take2/_search -d '@take2/keys.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take2/_search -d '@take2/orange_sizes.json'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:9200/take2/_search -d '@take2/orange_xl_facet.json'
```