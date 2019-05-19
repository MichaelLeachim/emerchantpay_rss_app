# This is a second part of the RSS reader task for the Emerchantpay

The goal is to:

Create an app that must require the Rss Reader package. It should parse multiple rss
feeds, store them in a JSON file and print them on the screen.

## How to install run and test

```bash

# Install
git clone https://github.com/MichaelLeachim/emerchantpay_rss_app;
cd emerchantpay_rss_app; 

# Run tests
go test; 

```

## How to use

```bash

### In order to save as json
./emerchantpay_rss_app --out=testdata/out.json -save -urlset="https://news.ycombinator.com/rss https://www.theguardian.com/uk/rss"

### In order to display
./emerchantpay_rss_app --in=testdata/out.json -print 

```
