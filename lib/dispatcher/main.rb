# frozen_string_literal: true

require_relative './telegram.rb'
require_relative './github.rb'

require 'erb'
require 'yaml'

github = GithubHandler.new(ENV['GITHUB_TOKEN'])

telegram = TelegramHandler.new(
  ENV['TELEGRAM_TOKEN'],
  ENV['TELEGRAM_CHATID']
)

goweekly, sha = github.get_file(
  ENV['REPOSITORY'],
  './content/goweekly.yaml'
)

github_update = false
goweekly.map do |content|
  next if content['send']

  content['send'] = true
  github_update = true

  template = ERB.new(
    File.read('template/goweekly.md.erb')
  )

  telegram.send_message(
    template.result_with_hash(
      content: content['content'],
      name: content['name']
    )
  )
end

if github_update
  puts('Running goweekly updates github')
  github.update_file(
    ENV['REPOSITORY'],
    './content/goweekly.yaml',
    'update goweekly data',
    goweekly.to_yaml.gsub("---\n", ''),
    sha
  )
else
  puts('no update')
end
