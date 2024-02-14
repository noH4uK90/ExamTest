//
//  TestView.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import SwiftUI

struct TestView: View {
    var body: some View {
        Content()
    }

    struct Content: View {

        @StateObject var viewModel: ViewModel

        init() {
            _viewModel = StateObject(wrappedValue: ViewModel())
        }

        var body: some View {
            EmptyView()
        }
    }
}

#Preview {
    TestView()
}
