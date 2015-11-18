require 'certifi'
require 'vkontakte_api'

require 'vnehm/token_manager'

module Vnehm

  ##
  # Client module contains all VK API interaction methods

  module Client

    ##
    # VK API client ID

    CLIENT_ID = 5144754

    ##
    # VK API client object

    VkontakteApi.configure do |config|
      config.log_requests = false
      config.log_errors = false
    end

    VK_CLIENT = VkontakteApi::Client.new(TokenManager.token)

    ##
    # SSL certificate file path

    ENV['SSL_CERT_FILE'] = Certifi.where

    def self.authorization_url
      VkontakteApi.authorization_url(type: :client,
                                     scope: [:audio, :offline],
                                     client_id: CLIENT_ID,
                                     redirect_uri: 'http://api.vkontakte.ru/blank.html')
    end

    ##
    # Returns raw array of likes or posts (depends on argument 'type')

    def self.tracks(count, offset)
      VK_CLIENT.audio.get(count: count, offset: offset)
    end

    def self.search(query, limit, offset)
      VK_CLIENT.audio.search(q: query,
                             count: limit,
                             offset: offset)
    end

  end
end
