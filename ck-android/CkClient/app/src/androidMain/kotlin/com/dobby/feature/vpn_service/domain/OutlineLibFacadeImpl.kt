package com.dobby.feature.vpn_service.domain

import kotlin_exports.OutlineDevice
import com.dobby.feature.vpn_service.OutlineLibFacade
import kotlin_exports.Kotlin_exports

internal class OutlineLibFacadeImpl: OutlineLibFacade {

    private var device: OutlineDevice? = null

    override fun init(apiKey: String) {
        device = Kotlin_exports.newOutlineDevice(apiKey)
    }

    override fun disconnect() {
        TODO("Not yet implemented")
    }

    override fun writeData(data: ByteArray) {
        device?.write(data)
    }

    override fun readData(): ByteArray? {
        return device?.read()
    }
}
