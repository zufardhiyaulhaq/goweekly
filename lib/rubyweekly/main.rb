require_relative './goweekly.rb'
require_relative './github.rb'

require 'yaml'

github = GithubHandler.new(ENV['GITHUB_TOKEN'])
goweekly = GoweeklyHandler.new

goweekly_new = goweekly.scrape_information
goweekly_old, sha = github.get_file(
  ENV['REPOSITORY'],
  './content/goweekly.yaml'
)

update = true
goweekly_old.each do |content|
  if goweekly_new['name'] == content['name']
    update = false
  end
end

if update
  puts('Running goweekly updates')
  goweekly_old.push(goweekly_new)
  github.update_file(
    ENV['REPOSITORY'],
    './content/goweekly.yaml',
    'update goweekly data',
    goweekly_old.to_yaml.gsub("---\n", ''),
    sha
  )
else
  puts('No goweekly updates')
end
