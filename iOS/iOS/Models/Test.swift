//
//  Test.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation

struct Test: Codable, Identifiable {
    let id: Int
    let name: String
    let questions: [Question]
}
