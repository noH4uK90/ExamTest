//
//  TestViewModel.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation
import Combine

extension TestView {
    @MainActor class ViewModel: ObservableObject {
        @Published var tests = [Test]()

        private var bag = Set<AnyCancellable>()
        @Inject private var networkService: NetworkProtocol

        init() {
            fetchTests()
        }

        func fetchTests() {
            Task {
                try networkService.getTests()
                    .receive(on: RunLoop.main)
                    .sink(
                        receiveCompletion: { _ in },
                        receiveValue: { [weak self] value in
                            self?.tests = value
                        }
                    )
                    .store(in: &bag)
            }
        }
    }
}
