module Vnehm

  # Primitive for SoundCloud track

  class Track

    attr_reader :hash

    def initialize(hash)
      @hash = hash
    end

    def artist
      CGI::unescapeHTML(@hash.artist)
    end

    def duration
      seconds = @hash['duration']

      time = Time.at(seconds)
      time -= time.utc_offset

      time.hour > 0 ? time.strftime("%H:%M:%S") : time.strftime("%M:%S")
    end

    def file_name
      "#{full_name.tr(",./'\\\"$%", '')}.mp3"
    end

    def file_path
      File.join(ENV['dl_path'], file_name)
    end

    def full_name
      "#{artist} - #{title}"
    end

    def id
      @hash.aid
    end

    def title
      CGI::unescapeHTML(@hash.title)
    end

    def url
      @hash.url
    end

  end
end
