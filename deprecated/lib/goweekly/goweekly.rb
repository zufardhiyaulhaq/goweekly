require 'active_support/core_ext/hash/keys'
require 'httparty'
require 'nokogiri'

class GoweeklyHandler
  def initialize
    @url = 'https://golangweekly.com/issues/latest?layout=bare'
  end

  private
  def parse_url(url)
    unparsed_page = HTTParty.get(url)
    Nokogiri::HTML(unparsed_page)
  end

  private
  def scrape_content
    content = []
    parsed_page = parse_url(@url)
    data = parsed_page.css('div#content')

    data.css('span.mainlink').each do |i|
      map = {
        'title' => i.text,
        'url' => i.css('a').attr('href').value
      }
      content.push(map)
    end

    content
  end

  public
  def scrape_information
    parsed_page = parse_url(@url)
    data = parsed_page.css('div#content')
    content = scrape_content

    information = {
      'name' => 'goweekly ' + data.css('table').text.match(/#\d+/).to_s,
      'send' => false,
      'content' => content
    }
  end
end
