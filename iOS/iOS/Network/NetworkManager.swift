//
//  NetworkManager.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation

enum Endpoints {
    case tests

    var baseURL: URL { URL(string: "http://localhost:8082/api")! }

    func path() -> String {
        switch self {
        case .tests:
            return "/test"
        }
    }

    var absoluteURL: URL? {
        let queryURL = baseURL.appending(path: self.path())
        let components = URLComponents(url: queryURL, resolvingAgainstBaseURL: true)
        guard var urlComponents = components else {
            return nil
        }

        switch self {
        case .tests:
            urlComponents.queryItems = []
        }

        return urlComponents.url
    }
}
