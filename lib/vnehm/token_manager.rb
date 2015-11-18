require 'cgi'

module Vnehm

  module TokenManager

    def self.token
      Cfg['token'] if Cfg.exist?
    end

    def self.authorize
      auth_url = Client.authorization_url
      UI.say "1. Переходите по этой ссылке с помощью вашего браузера: #{auth_url.magenta}"
      UI.say "2. Жмите #{'Разрешить'.green}"
      UI.say '3. Скопируйте ссылку из адресной строки вашего браузера сюда:'
      uri = UI.ask

      uri = URI(uri)
      hash = CGI::parse(uri.fragment)

      if hash['access_token']
        Cfg['token'] = hash['access_token'].first
      else
        UI.term 'Вы ввели неверную ссылку! ' \
                'Пожалуйста, авторизируйтесь ещё раз и перепроверьте введеные данные'
      end

      UI.success 'Вы успешно авторизованы!'
    end

  end

end
