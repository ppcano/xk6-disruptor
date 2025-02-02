# xk6-disruptor

</br>
</br>

<div align="center">

![logo](assets/logo.png)


</div>

The xk6-disruptor is a [k6](https://k6.io) extension providing fault injection capabilities to test system's reliability under turbulent conditions. Think of it as "like unit testing, but for reliability". 

This project aims to aid developers in building reliable systems, implementing the goals of "Chaos Engineering" discipline in a k6 way - with the best developer experience as its primary objective. 

xk6-disruptor is intended for systems running in kubernetes. Other platforms are not supported at this time.

The extension offers an [API](https://k6.io/docs/javascript-api/xk6-disruptor/api) for creating disruptors that target one specific type of component (for example, Pods) and are capable of injecting different types of [faults](https://k6.io/docs/javascript-api/xk6-disruptor/api/faults) such as errors in HTTP requests served by that component. Currently disruptors exist for [Pods](https://k6.io/docs/javascript-api/xk6-disruptor/api/poddisruptor) and [Services](https://k6.io/docs/javascript-api/xk6-disruptor/api/servicedisruptor), but others will be introduced in the future as well as additional types of faults for the existing disruptors.

> ⚠️  xk6-disruptor is in the alpha stage, undergoing active development. We do not guarantee API compatibility between releases - your k6 scripts may need to be updated on each release until this extension reaches v1.0 release.

If you encounter any bugs or unexpected behavior, please search the [currently open GitHub issues](https://github.com/grafana/xk6-disruptor/issues) first, and create a new one if it doesn't exist yet.

## Use case for xk6-disruptor

The main use case for xk6-disruptor is to test the resiliency of an application of diverse types disruptions by reproducing their effects, but without having to reproduce their root-causes. For example, inject delays in the HTTP requests an application makes to a service  without having to stress or interfere with the infrastructure (network, nodes) on which the service runs, or affecting other workloads in unexpected ways.

In this way, xk6-disruptor make reliability tests repeatable, predictable and limits their blast-radius. These are important characteristic in order to incorporate this kind of tests in the test suits of applications deployed on shared infrastructures such as staging environments.

## Learn more

Check the [get started guide](https://k6.io/docs/javascript-api/xk6-disruptor/getstarted) for instructions on how to install and use `xk6-disruptor`.

The [Roadmap](/ROADMAP.md) presents the project's goals for the coming months regarding new functionalities and enhancements.

If you are interested in contributing with the development of this project, check the [contributing guide](/docs/01-development/01-contributing.md)
