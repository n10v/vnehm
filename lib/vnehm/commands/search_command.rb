require 'vnehm/tracks_view_command'

module Vnehm

  class SearchCommand < TracksViewCommand

    def initialize
      super

      add_option(:"-t", '-t ПУТЬ',
                 'Скачивать аудиозапись(и) в ПУТЬ')

      add_option(:"-pl", '-pl ПЛЕЙЛИСТ',
                 "Добавлять аудиозапись(и) в плейлист iTunes'a c именем" \
                                                                  ' ПЛЕЙЛИСТ')

      add_option(:"-lim", '-lim ЧИСЛО',
                 'Показывать ЧИСЛО аудиозаписей на каждой страницe')

      add_option(:"-dl", '-dl yes',
                 'Не добавлять аудиозаписи в iTunes. Просто скачать их')

    end

    def execute
      # Convert dash-options to normal options
      options_to_convert = { :"-t"   => :to,
                             :"-pl"  => :pl,
                             :"-lim" => :limit,
                             :"-dl"  => :dl }

      options_to_convert.each do |k,v|
        value = @options[k]
        @options.delete(k)
        @options[v] = value unless value.nil?
      end

      @query = @options[:args].join(' ')
      super
    end

    def arguments
      { 'ЗАПРОС' => 'Искать аудиозаписи по ЗАПРОСу' }
    end

    def program_name
      'vnehm search'
    end

    def summary
      'Поиск, вывод и загрузка определенных аудиозаписей по запросу'
    end

    def usage
      "#{program_name} ЗАПРОС [ОПЦИИ]"
    end

    protected

    def get_tracks
      UI.term 'Вы должны ввести запрос' if @query.empty?

      @track_manager.search(@query, @limit, @offset)
    end

  end
end
