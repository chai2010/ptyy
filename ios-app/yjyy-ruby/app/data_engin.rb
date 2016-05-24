# Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

class DataEngin
  def self.Search(query, limits)
    result = []
    YjyySearch(query, limits).each_line do |line|
      result << line
    end
    result
  end
end
