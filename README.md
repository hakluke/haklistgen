# haklistgen
Turns any junk text into a usable wordlist for brute-forcing.

# Installation

```
go get -u github.com/hakluke/haklistgen
```

# Usage Examples

Scrape all words out of an HTTP response to build a directory bruteforce wordlist:

```
curl https://wikipedia.org | haklistgen
```

Pipe a list of subdomains to it to generate a wordlist for bruteforcing more subdomains:

```
subfinder -silent -d example.com | haklistgen
```

Piping in a custom JavaScript file could yield some interesting results:

```
curl https://example.com/app.js | haklistgen
```

You could create a great custom wordlist for a large-scope target doing something like this:

```
subfinder -silent -d hakluke.com | anew subdomains.txt | httpx -silent | anew urls.txt | hakrawler | anew endpoints.txt | while read url; do curl $url --insecure | haklistgen | anew wordlist.txt; done
cat subdomains.txt urls.txt endpoints.txt | haklistgen | anew wordlist.txt;
```

This would save subdomains to `subdomains.txt`, then save httpx output to `urls.txt`, then crawl each url and save the hakrawler output to `endpoints.txt`, then fetch every URL in `endpoints.txt` and make a wordlist out of it, concatenating all of the wordlists to `wordlist.txt`. Then it takes all of the subdomains and urls, and adds words out of the words in those too.
