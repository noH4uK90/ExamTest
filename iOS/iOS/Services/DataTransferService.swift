//
//  DataTransferService.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation
import Combine

protocol DataTransferProtocol {
    func fetch<T: Codable>(_ url: URL, _ model: T.Type) -> AnyPublisher<T, Error>

    func post<TData: Codable, TResult: Codable>(_ url: URL, _ data: TData) throws -> AnyPublisher<TResult, Error>
}

class DataTransferService: DataTransferProtocol {
    private var decoder = JSONDecoder()
    private var encoder = JSONEncoder()

    func fetch<T>(_ url: URL, _ model: T.Type) -> AnyPublisher<T, Error> where T : Decodable, T : Encodable {
        return URLSession.shared.dataTaskPublisher(for: url)
            .map({ $0.data })
            .decode(type: T.self, decoder: decoder)
            .receive(on: DispatchQueue.main)
            .eraseToAnyPublisher()
    }

    func post<TData, TResult>(_ url: URL, _ data: TData) throws -> AnyPublisher<TResult, Error> where TData : Decodable, TData : Encodable, TResult : Decodable, TResult : Encodable {
        let jsonData = try encoder.encode(data)

        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        request.httpBody = jsonData

        return URLSession.shared.dataTaskPublisher(for: request)
            .map({ $0.data })
            .decode(type: TResult.self, decoder: decoder)
            .receive(on: DispatchQueue.main)
            .eraseToAnyPublisher()
    }
}
